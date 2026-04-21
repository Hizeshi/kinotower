package servise

import (
	"context"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/genders/repository"
)

type GenderService struct {
	repo *repository.GenderRepository
}

func NewGenderService(repo *repository.GenderRepository) *GenderService {
	return &GenderService{repo: repo}
}

func (s *GenderService) GetAllGenders(ctx context.Context) ([]domain.Gender, error) {
	return s.repo.GetAll(ctx)
}