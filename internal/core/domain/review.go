package domain

type ReviewUser struct {
	ID  int    `json:"id"`
	FIO string `json:"fio"`
}

type FilmReview struct {
	ID        int        `json:"id"`
	User      ReviewUser `json:"user"`
	Message   string     `json:"message"`
	CreatedAt string     `json:"created_at"`
}

type ReviewFilm struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserReviewResponse struct {
	ID         int        `json:"id"`
	Film       ReviewFilm `json:"film"`
	Message    string     `json:"message"`
	IsApproved int        `json:"is_approved"` 
	CreatedAt  string     `json:"created_at"`
}

type CreateReviewRequest struct {
	FilmID  int    `json:"film_id"`
	Message string `json:"message"`
}