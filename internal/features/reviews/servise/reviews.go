package servise

import (
	"context"
	"errors"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/reviews/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) AddReview(ctx context.Context, userID int, req domain.CreateReviewRequest) (domain.UserReviewResponse, error) {
	
	reviewID, err := s.repo.CreateReview(ctx, userID, req)
	if err != nil {
		return domain.UserReviewResponse{}, err
	}

	return s.repo.GetReviewByID(ctx, reviewID)
}

func (s *ReviewService) GetUserReviews(ctx context.Context, userID int) ([]domain.UserReviewResponse, error) {

	return s.repo.GetUserReviews(ctx, userID)
}

func (s *ReviewService) DeleteReview(ctx context.Context, userID int, reviewID int) error {
	_, err := s.repo.GetReviewByID(ctx, reviewID)
	if err != nil {
		return errors.New("Review not found") 
	}

	return s.repo.DeleteReview(ctx, reviewID)
}