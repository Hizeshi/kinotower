package domain

type FilmCountry struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type FilmCategory struct {
    ID int `json:"id"`
	Name string `json:"name"`
}

type FilmFilter struct {
	Page int
	Size int
	SortBy string
	SortDir string
	Category int
	Country int
	Search string
}

type Film struct {
     ID int `json:"id"`
	 Name string `json:"name"`
	 Durarion int `json:"duration"`
	 YearOfIssue int `json:"year_of_issue"`
	 Age int `json:"age"`
	 LinkImg string `json:"link_img"`
	 LinkKinopoisk string `json:"link_kinopoisk"`
	 LinkVideo string `json:"link_video"`
	 CreatedAt string `json:"created_at"`

	 Country FilmCountry `json:"country"`
	 Categories []FilmCategory `json:"categories"`

	 RatingAvg float64 `json:"ratingAvg"`
	 ReviewCount int `json:"reviewCount"`
}

type FilmResponse struct {
	Page int `json:"page"`
	Size int `json:"size"`
	Total int `json:"total"`
    Films []Film `json:"films"`
}