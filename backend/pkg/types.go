package pkg

// UserDTO represents the public user data (no password)
// This is shared across domains to avoid import cycles
type UserDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// LoginRequest represents the login payload
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserRequest represents the user creation payload
type CreateUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
