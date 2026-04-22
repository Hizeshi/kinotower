package repository

import (
	"context"
	"fmt"
	"strings"

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
