package user

import "digitalwallet/backend/pkg"

// User represents the internal user entity with sensitive data
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"` // hashed password
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// ToDTO converts User to UserDTO (removes sensitive data)
func (u *User) ToDTO() pkg.UserDTO {
	return pkg.UserDTO{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
