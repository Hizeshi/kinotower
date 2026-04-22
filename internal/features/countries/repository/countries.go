package repository

import (
	"context"

	"github.com/nedpals/supabase-go"
	"github.com/Hizeshi/kinotower/internal/core/domain"
)


type CountryRepository struct {
	db *supabase.Client
}

func NewCountryRepository(db *supabase.Client) *CountryRepository {
	return &CountryRepository{db: db}
}

func (r *CountryRepository) GetAll(ctx context.Context) ([]domain.Country, error) {
	var countries []domain.Country

	err := r.db.DB.From("countries_view").Select("*").Execute(&countries)
	if err != nil {
		return nil, err
	}

	if countries == nil {
		countries = []domain.Country{}
	}

	return countries, nil
}