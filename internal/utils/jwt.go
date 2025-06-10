package utils

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo-api/config"
	"todo-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents the claims in the JWT
type JWTClaims struct {
	UserID uint   `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user models.User) (string, error) {
	// Get JWT secret from config
	jwtSecret := config.GetConfig().JWTSecret
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not configured")
	}

	// Set token expiration time (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create claims with user information
	claims := &JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "todo-api",
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*JWTClaims, error) {
	// Get JWT secret from config
	jwtSecret := config.GetConfig().JWTSecret
	if jwtSecret == "" {
		return nil, errors.New("JWT secret is not configured")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	// Extract and return claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// Context key for user ID
type contextKey string

const userIDKey contextKey = "userID"

// SetUserIDInContext adds the user ID to the context
func SetUserIDInContext(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext retrieves the user ID from the context
func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(userIDKey).(uint)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userID, nil
}
