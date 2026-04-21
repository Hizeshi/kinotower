package repository

import (
	"context"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/nedpals/supabase-go"
)

type GenderRepository struct {
	db *supabase.Client
}

func NewGenderRepository(db *supabase.Client) *GenderRepository {
	return &GenderRepository{db: db}
}

func (r *GenderRepository) GetAll(ctx context.Context) ([]domain.Gender, error) {

	var genders []domain.Gender

	err := r.db.DB.From("genders").Select("*").Execute(&genders)
	
	if err != nil {
		return nil, err
	}

	return genders, nil
}