package servise

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/auth/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("super_secret_kinotower_key")

type AuthServise struct {
	repo *repository.AuthRepository
}

func (s *AuthServise) RevokeToken(ctx context.Context, token string, expiresAt time.Time) error {
    return s.repo.AddToBlacklist(ctx, token, expiresAt)
}

func NewAuthServise(repo *repository.AuthRepository) *AuthServise {
	return &AuthServise{repo: repo}
}

func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtSecretKey)
}

func (s *AuthServise) SignUp(ctx context.Context, req domain.SignUpRequest) (domain.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.AuthResponse{}, err
	}

	newUser := domain.User{
		Fio:      req.Fio,
		Email:    req.Email,
		Password: string(hashedPassword),
		Birthday: req.Birthday,
		GenderID: req.GenderID,
	}

	createdUser, err := s.repo.CreateUser(ctx, newUser)
	if err != nil {
		slog.Error("Failed to create user", "error", err)
		return domain.AuthResponse{}, errors.New("email already exists or invalid data")
	}

	token, err := generateToken(createdUser.ID)
	if err != nil {
		slog.Error("Failed to generate token", "error", err)
		return domain.AuthResponse{}, err
	}

	return domain.AuthResponse{
		Status: "success",
		Token:  token,
		ID:     createdUser.ID,
		Fio:    createdUser.Fio,
	}, nil
}

func (s *AuthServise) SignIn(ctx context.Context, req domain.SignInRequest) (domain.AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		slog.Error("SignIn DB Error: failed to get user", "error", err, "email", req.Email)
		return domain.AuthResponse{}, errors.New("Wrong email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		slog.Error("SignIn Bcrypt Error: password mismatch", "error", err)
		return domain.AuthResponse{}, errors.New("Wrong email or password")
	}

	token, err := generateToken(user.ID)
	if err != nil {
		slog.Error("Failed to generate token", "error", err)
		return domain.AuthResponse{}, err
	}

	return domain.AuthResponse{
		Status: "success",
		Token:  token,
		ID:     user.ID,
		Fio:    user.Fio,
	}, nil
}


