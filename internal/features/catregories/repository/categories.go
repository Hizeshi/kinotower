package repository

import (
	"context"

	"github.com/nedpals/supabase-go"
	"github.com/Hizeshi/kinotower/internal/core/domain"
)

type CategoryRepository struct {
	db *supabase.Client
}

func NewCategoryRepository(db *supabase.Client) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category

	err := r.db.DB.From("categories_view").Select("*").Execute(&categories)
	if err != nil {
		return nil, err
	}

	if categories == nil {
		categories = []domain.Category{}
	}

	return categories, nil
}