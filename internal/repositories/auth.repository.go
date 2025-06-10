package repositories

import (
	"context"
	"errors"
	"net/http"
	"todo-api/database"
	"todo-api/internal/dtos"
	"todo-api/internal/models"
	"todo-api/internal/utils"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewAuthRepository(logger *zap.Logger) *AuthRepository {
	return &AuthRepository{DB: database.GetDB(), Logger: logger}
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerUserDto dtos.RegisterUserDto) (dtos.StructuredResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserDto.Password), bcrypt.DefaultCost)

	if err != nil {
		r.Logger.Error("Failed to hash password", zap.Error(err))
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to register user",
			Payload: nil,
		}, err
	}

	user := models.User{
		Email:        registerUserDto.Email,
		Name:         registerUserDto.Name,
		PasswordHash: string(hashedPassword),
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to register user",
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "User registered successfully",
		Payload: map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	}, nil
}

func (r *AuthRepository) LoginUser(ctx context.Context, loginUserDto dtos.LoginUserDto) (dtos.StructuredResponse, error) {
	var user models.User

	// Find the user by email
	if err := r.DB.Where("email = ?", loginUserDto.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.StructuredResponse{
				Success: false,
				Status:  http.StatusUnauthorized,
				Message: "Invalid email or password",
				Payload: nil,
			}, nil
		}
		r.Logger.Error("Failed to find user", zap.Error(err))
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to login",
			Payload: nil,
		}, err
	}

	// Compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginUserDto.Password))
	if err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Message: "Invalid email or password",
			Payload: nil,
		}, nil
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		r.Logger.Error("Failed to generate token", zap.Error(err))
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to login",
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Login successful",
		Payload: map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"token": token,
		},
	}, nil
}
