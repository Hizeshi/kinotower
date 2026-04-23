package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/nedpals/supabase-go"
)

type AuthRepository struct {
	db *supabase.Client
}

func NewAuthRepository(db *supabase.Client) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context,email string) (domain.User, error) {
	var users []domain.User

	cleanEmail := strings.TrimSpace(email)

	err := r.db.DB.From("users").Select("*").Filter("email", "ilike", cleanEmail).Execute(&users)
	if err != nil {
		return domain.User{}, err
	}

	if len(users) == 0 {
		return domain.User{}, fmt.Errorf("user not found")
	}

	return users[0], nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	var results []domain.User

	err := r.db.DB.From("users").Insert(user).Execute(&results)
	if err != nil {
		return domain.User{}, err
	}

	if len(results) == 0 {
		return domain.User{}, fmt.Errorf("failed to insert user")
	}

	return results[0], nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (r *AuthRepository) AddToBlacklist(ctx context.Context, token string, expiresAt time.Time) error {
	data := map[string]interface{}{
		"token":      hashToken(token),
		"expires_at": expiresAt.Format(time.RFC3339),
	}
	var results []map[string]interface{}
	err := r.db.DB.From("blacklisted_tokens").Insert(data).Execute(&results)
	if err != nil {
		return fmt.Errorf("failed to insert token: %w", err)
	}
	return nil
}

func (r *AuthRepository) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	var results []map[string]interface{}
	err := r.db.DB.From("blacklisted_tokens").Select("id").Eq("token", hashToken(token)).Execute(&results)
	if err != nil {
		return false, err
	}
	return len(results) > 0, nil
}
