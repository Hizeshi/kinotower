package servise

import (
	"context"
	"errors"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/ratings/repository"
)

var ErrScoreExist = errors.New("Score exist")

type RatingService struct {
	repo *repository.RatingRepository
}

func NewRatingService(repo *repository.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) AddRating(ctx context.Context, userID int, req domain.CreateRatingRequest) (int, error) {
	exists, _ := s.repo.CheckExist(ctx, userID, req.FilmID)
	if exists {
		return 0, ErrScoreExist
	}

	return s.repo.Create(ctx, userID, req)
}

func (s *RatingService) GetUserRatings(ctx context.Context, userID int) ([]domain.RatingResponse, error) {
	return s.repo.GetUserRatings(ctx, userID)
}

func (s *RatingService) DeleteRating(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
