package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/nedpals/supabase-go"
)

type UserRepository struct {
	db *supabase.Client
}

func NewUserRepository(db *supabase.Client) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserProfileByID(ctx context.Context, id int) (domain.UserProfile, error) {
	var profiles []domain.UserProfile

	err := r.db.DB.From("user_profile_view").Select("*").Eq("id", fmt.Sprintf("%d", id)).Execute(&profiles)
	if err != nil {
		return domain.UserProfile{}, err
	}

	if len(profiles) == 0 {
		return domain.UserProfile{}, fmt.Errorf("user not found")
	}

	return profiles[0], nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user domain.User) error {
	data := map[string]interface{}{
		"fio":       user.Fio,
		"email":     user.Email,
		"birthday":  user.Birthday,
		"gender_id": user.GenderID,
	}

	err := r.db.DB.From("users").Update(data).Eq("id", fmt.Sprintf("%d", id)).Execute(nil)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	data := map[string]interface{}{
		"deleted_at": time.Now().Format(time.RFC3339),
	}

	err := r.db.DB.From("users").Update(data).Eq("id", fmt.Sprintf("%d", id)).Execute(nil)
	return err
}
