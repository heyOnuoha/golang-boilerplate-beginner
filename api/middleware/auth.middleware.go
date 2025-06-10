package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-api/internal/dtos"
	"todo-api/internal/utils"

	"go.uber.org/zap"
)

// AuthMiddleware is a middleware that checks if the request has a valid JWT token
func AuthMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("Missing Authorization header")
				respondWithError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Check if the Authorization header has the Bearer prefix
			if !strings.HasPrefix(authHeader, "Bearer ") {
				logger.Warn("Invalid Authorization header format")
				respondWithError(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate the token
			claims, err := utils.ValidateToken(tokenString)
			if err != nil {
				logger.Warn("Invalid token", zap.Error(err))
				respondWithError(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Add the user ID to the request context
			ctx := r.Context()
			ctx = utils.SetUserIDInContext(ctx, claims.UserID)
			r = r.WithContext(ctx)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Helper function to respond with an error
func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := dtos.StructuredResponse{
		Success: false,
		Status:  statusCode,
		Message: message,
		Payload: nil,
	}
	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If encoding fails, log the error and write a simple error message
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
	}
}
