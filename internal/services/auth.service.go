package services

import (
	"context"
	"todo-api/internal/dtos"
	"todo-api/internal/repositories"

	"go.uber.org/zap"
)

type AuthService struct {
	logger *zap.Logger
	repo   *repositories.AuthRepository
}

func NewAuthService(logger *zap.Logger) *AuthService {
	return &AuthService{
		logger: logger,
		repo:   repositories.NewAuthRepository(logger),
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, registerUserDto dtos.RegisterUserDto) (dtos.StructuredResponse, error) {
	return s.repo.RegisterUser(ctx, registerUserDto)
}

func (s *AuthService) LoginUser(ctx context.Context, loginUserDto dtos.LoginUserDto) (dtos.StructuredResponse, error) {
	return s.repo.LoginUser(ctx, loginUserDto)
}
