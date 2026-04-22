package servise

import (
	"context"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/films/repository"
)

type FilmService struct {
	repo *repository.FilmRepository
}

func NewFilmService(repo *repository.FilmRepository) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) GetFilms(ctx context.Context, filter domain.FilmFilter) (domain.FilmResponse, error) {

	if filter.Page <= 0 {
		filter.Page = 1 
	}
	if filter.Size <= 0 {
		filter.Size = 10 
	}
	if filter.SortBy == "" {
		filter.SortBy = "name" 
	}
	if filter.SortDir != "desc" {
		filter.SortDir = "asc" 
	}

	films, total, err := s.repo.GetFilms(ctx, filter)
	if err != nil {
		return domain.FilmResponse{}, err
	}

	response := domain.FilmResponse{
		Page:  filter.Page,
		Size:  len(films), 
		Total: total,
		Films: films,
	}

	return response, nil
}

func (s *FilmService) GetFilmByID(ctx context.Context, id int) (domain.Film, error) {
	return s.repo.GetFilmByID(ctx, id)
}

func (s *FilmService) GetFilmReviews(ctx context.Context, filmID int) ([]domain.FilmReview, error) {
	
	_, err := s.GetFilmByID(ctx, filmID)
	if err != nil {
		return nil, err
	}

	return s.repo.GetReviewsByFilmID(ctx, filmID)
	
}
