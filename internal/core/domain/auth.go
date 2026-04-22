package domain

type SignUpRequest struct {
	Fio      string `json:"fio"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
	GenderID int    `json:"gender_id"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
	ID     int    `json:"id"`
	Fio    string `json:"fio"`
}

type User struct {
	ID        int    `json:"id,omitempty"` 
	Fio       string `json:"fio"`
	Birthday  string `json:"birthday"`
	GenderID  int    `json:"gender_id"`
	Email     string `json:"email"`
	Password  string `json:"password"` 
	CreatedAt string `json:"created_at,omitempty"`
}