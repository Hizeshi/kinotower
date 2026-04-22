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
	var allFilms []domain.Film

	dbSortColumn := filter.SortBy
	if dbSortColumn == "rating" {
		dbSortColumn = "ratingAvg"
	}

	selectQuery := r.db.DB.From("films_view").Select("*").
		OrderBy(dbSortColumn, filter.SortDir)

	query := selectQuery.Gt("id", "0")

	if filter.Search != "" {
		query = query.Filter("name", "ilike", "*"+filter.Search+"*")
	}

	if filter.Country > 0 {
		query = query.Eq("country_id", fmt.Sprintf("%d", filter.Country))
	}

	if filter.Category > 0 {
		query = query.Filter("category_ids", "cs", fmt.Sprintf("{%d}", filter.Category))
	}

	err := query.Execute(&allFilms)
	if err != nil {
		return nil, 0, err
	}

	total := len(allFilms)

	start := (filter.Page - 1) * filter.Size
	end := start + filter.Size

	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	pagedFilms := allFilms[start:end]

	if len(pagedFilms) == 0 {
		pagedFilms = []domain.Film{}
	}

	return pagedFilms, total, nil
}

func (r *FilmRepository) GetFilmByID(ctx context.Context, id int) (domain.Film, error) {
	var films []domain.Film

	err := r.db.DB.From("films_view").Select("*").Eq("id", fmt.Sprintf("%d", id)).Execute(&films)
	if err != nil {
		return domain.Film{}, err
	}

	if len(films) == 0 {
		return domain.Film{}, fmt.Errorf("film not found")
	}

	return films[0], nil
}

func (r *FilmRepository) GetReviewsByFilmID(ctx context.Context, filmID int) ([]domain.FilmReview, error) {
	var reviews []domain.FilmReview

	err := r.db.DB.From("film_reviews_view").Select("*").Eq("film_id", fmt.Sprintf("%d", filmID)).Execute(&reviews)
	if err != nil {
		return nil, err
	}

	if reviews == nil {
		reviews = []domain.FilmReview{}
	}

	return reviews, nil
}