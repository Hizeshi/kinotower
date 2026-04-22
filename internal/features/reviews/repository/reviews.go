package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/nedpals/supabase-go"
)

type ReviewRepository struct {
	db *supabase.Client
}

func NewReviewRepository(db *supabase.Client) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(ctx context.Context, userID int, req domain.CreateReviewRequest) (int, error) {
	var result []struct { ID int `json:"id"` }
	
	data := map[string]interface{}{
		"user_id":     userID,
		"film_id":     req.FilmID,
		"message":     req.Message,
		"is_approved": false, 
	}

	err := r.db.DB.From("reviews").Insert(data).Execute(&result)
	if err != nil || len(result) == 0 {
		return 0, err
	}
	return result[0].ID, nil
}

func (r *ReviewRepository) GetUserReviews(ctx context.Context, userID int) ([]domain.UserReviewResponse, error) {
	var reviews []domain.UserReviewResponse
	err := r.db.DB.From("user_reviews_view").Select("*").Eq("user_id", fmt.Sprintf("%d", userID)).Execute(&reviews)
	return reviews, err
}

func (r *ReviewRepository) GetReviewByID(ctx context.Context, reviewID int) (domain.UserReviewResponse, error) {
	var reviews []domain.UserReviewResponse

	err := r.db.DB.From("user_reviews_view").Select("*").Eq("id", fmt.Sprintf("%d", reviewID)).Execute(&reviews)
	if err != nil {
		return domain.UserReviewResponse{}, err
	}

	if len(reviews) == 0 {
		return domain.UserReviewResponse{}, fmt.Errorf("review not found")
	}

	return reviews[0], nil
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, reviewID int) error {
	data := map[string]interface{}{
		"deleted_at": time.Now().Format(time.RFC3339),
	}

	err := r.db.DB.From("reviews").Update(data).Eq("id", fmt.Sprintf("%d", reviewID)).Execute(nil)
	return err
}
