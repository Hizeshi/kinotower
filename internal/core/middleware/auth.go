package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("super_secret_kinotower_key")

type ContextKey string

const UserIDKey ContextKey = "userID"

type AuthRepository interface {
	IsBlacklisted(ctx context.Context, token string) (bool, error)
}

type AuthMiddleware struct {
	repo AuthRepository
}

func NewAuthMiddleware(repo AuthRepository) *AuthMiddleware {
	return &AuthMiddleware{repo: repo}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		isRevoked, err := m.repo.IsBlacklisted(r.Context(), tokenString)
		if err != nil {
			http.Error(w, "Internal security error", http.StatusInternalServerError)
			return
		}

		if isRevoked {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "unauthorized",
				"message": "Token has been revoked (Signed out)",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := int(claims["userID"].(float64))
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	})
}