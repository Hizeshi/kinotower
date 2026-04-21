package repository

import (
	"context"
	"fmt"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/nedpals/supabase-go"
)

type FilmRepository struct {
	db *supabase.Client
}

func NewFilmRepository(db *supabase.Client) *FilmRepository {
	return &FilmRepository{db: db}
}

func (r *FilmRepository) GetFilms(ctx context.Context, filter domain.FilmFilter) ([]domain.Film, int, error) {
	var films []domain.Film

	start := (filter.Page - 1) * filter.Size
	end := start + filter.Size - 1

	selectQuery := r.db.DB.From("films").Select("*").
		OrderBy(filter.SortBy, filter.SortDir).
		Range(start, end)

	query := selectQuery.Gt("id", "0")

	if filter.Search != "" {
		query = query.Filter("name", "ilike", "%"+filter.Search+"%")
	}

	if filter.Country > 0 {
		query = query.Eq("country_id", fmt.Sprintf("%d", filter.Country))
	}

	err := query.Execute(&films)
	if err != nil {
		return nil, 0, err
	}

	total := len(films)

	return films, total, nil
}