package dtos

// RegisterUserDto represents the data needed to register a new user
// @Description Registration data for creating a new user account
type RegisterUserDto struct {
	// User's email address
	// @example john.doe@example.com
	Email string `json:"email" binding:"required" example:"john.doe@example.com"`
	// User's password (min 8 characters)
	// @example SecureP@ssw0rd
	Password string `json:"password" binding:"required" example:"SecureP@ssw0rd"`
	// Confirmation of the password
	// @example SecureP@ssw0rd
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password" example:"SecureP@ssw0rd"`
	// User's full name
	// @example John Doe
	Name string `json:"name" binding:"required" example:"John Doe"`
}

// LoginUserDto represents the data needed to login a user
// @Description Login credentials for authenticating a user
type LoginUserDto struct {
	// User's email address
	// @example john.doe@example.com
	Email string `json:"email" binding:"required" example:"john.doe@example.com"`
	// User's password
	// @example SecureP@ssw0rd
	Password string `json:"password" binding:"required" example:"SecureP@ssw0rd"`
}
