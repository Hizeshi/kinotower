package servise

import (
	"context"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/users/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func (s *UserService) GetUserProfile(context context.Context, userID int) (any, error) {
	panic("unimplemented")
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(ctx context.Context, userID int) (domain.UserProfile, error) {
	return s.repo.GetUserProfileByID(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user domain.User) error {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}
