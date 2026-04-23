package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/auth/servise"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	service *servise.AuthServise
}

func NewAuthHandler(service *servise.AuthServise) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req domain.SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.SignUp(r.Context(), req)
	if err != nil {
		slog.Error("Error during signup", "error", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) 
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req domain.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.SignIn(r.Context(), req)
	if err != nil {
		slog.Warn("Failed login attempt", "email", req.Email)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "invalid",
			"message": "Wrong email or password",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

    token, _, _ := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
    claims, _ := token.Claims.(jwt.MapClaims)
    exp := time.Unix(int64(claims["exp"].(float64)), 0)

    err := h.service.RevokeToken(r.Context(), tokenString, exp)
    if err != nil {
        http.Error(w, "Failed to sign out", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}