package domain

type ParentCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	ParentCategory *ParentCategory `json:"parentCategory"` 
	FilmCount      int             `json:"filmCount"`
}