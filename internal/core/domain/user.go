package domain

type UserGender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserProfile struct {
	ID          int        `json:"id"`
	Fio         string     `json:"fio"`
	Email       string     `json:"email"`
	Birthday    string     `json:"birthday"`
	Gender      UserGender `json:"gender"`
	ReviewCount int        `json:"reviewCount"`
	RatingCount int        `json:"ratingCount"`
}