package servise

import (
	"context"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/countries/repository"
)


type CountryService struct {
	repo *repository.CountryRepository
}

func NewCountryService(repo *repository.CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (s *CountryService) GetAllCountries(ctx context.Context) ([]domain.Country, error) {
	return s.repo.GetAll(ctx)
}