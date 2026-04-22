package repository 

import (
	"context"
	"fmt"
	"time"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/nedpals/supabase-go"
)

type RatingRepository struct {
	db *supabase.Client
}

func NewRatingRepository(db *supabase.Client) *RatingRepository {
	return &RatingRepository{db: db}
}

func (r *RatingRepository) CheckExist(ctx context.Context, userID, filmID int) (bool, error) {
	var results []interface{}
	err := r.db.DB.From("ratings").Select("id").
		Eq("user_id", fmt.Sprintf("%d", userID)).
		Eq("film_id", fmt.Sprintf("%d", filmID)).
		Is("deleted_at", "null").
		Execute(&results)
	return len(results) > 0, err
}

func (r *RatingRepository) Create(ctx context.Context, userID int, req domain.CreateRatingRequest) (int, error) {
	var result []struct { ID int `json:"id"` }
	
	data := map[string]interface{}{
		"user_id": userID,
		"film_id": req.FilmID,
		"ball":    req.Ball, 
	}

	err := r.db.DB.From("ratings").Insert(data).Execute(&result)
	if err != nil || len(result) == 0 {
		return 0, err
	}
	return result[0].ID, nil
}

func (r *RatingRepository) GetUserRatings(ctx context.Context, userID int) ([]domain.RatingResponse, error) {
	var ratings []domain.RatingResponse
	err := r.db.DB.From("user_ratings_view").Select("*").Eq("user_id", fmt.Sprintf("%d", userID)).Execute(&ratings)
	return ratings, err
}

func (r *RatingRepository) Delete(ctx context.Context, id int) error {
	data := map[string]interface{}{
		"deleted_at": time.Now().Format(time.RFC3339),
	}
	return r.db.DB.From("ratings").Update(data).Eq("id", fmt.Sprintf("%d", id)).Execute(nil)
}