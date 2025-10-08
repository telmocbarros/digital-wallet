package user

// User represents the internal user entity with sensitive data
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"` // hashed password
}

// UserDTO represents the public user data (no password)
type UserDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// LoginRequest represents the login payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserRequest represents the user creation payload
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ToDTO converts User to UserDTO (removes sensitive data)
func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:    u.ID,
		Email: u.Email,
	}
}
