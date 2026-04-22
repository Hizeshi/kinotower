package domain

type RatingResponse struct {
	ID        int        `json:"id"`
	Film      ReviewFilm `json:"film"` 
	Score     int        `json:"score"`
	CreatedAt string     `json:"created_at"`
}

type CreateRatingRequest struct {
	FilmID int `json:"film_id"`
	Ball  int `json:"ball"`
}