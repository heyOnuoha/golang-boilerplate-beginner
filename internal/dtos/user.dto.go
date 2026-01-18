package dtos

// UserDto represents a user in the system
// @Description A user with all their details
type UserDto struct {
	// Unique identifier
	// @example 1
	ID uint `json:"id" example:"1"`
	// Email address
	// @example john.doe@example.com
	Email string `json:"email" example:"john.doe@example.com"`
	// Name of the user
	// @example John Doe
	Name string `json:"name" example:"John Doe"`
}
