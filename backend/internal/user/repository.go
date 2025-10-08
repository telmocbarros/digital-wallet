package user

import (
	"digitalwallet/backend/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Repository defines the interface for user data access
type Repository interface {
	GetByEmail(email string) (*User, error)
	GetByID(id string) (*User, error)
	GetAll() ([]UserDTO, error)
	Create(email, password string) (*User, error)
	VerifyCredentials(email, password string) (*UserDTO, error)
}

// inMemoryRepository implements Repository using in-memory storage
type inMemoryRepository struct {
	users []User
}

// NewRepository creates a new user repository with test data
func NewRepository() Repository {
	return &inMemoryRepository{
		users: []User{
			{ID: "b18b851a-c8c4-4957-b68a-14362a1810c6", Email: "john@example.com", Password: "$2a$14$4Il8GoD6jpuFDi4ScOAqWuRZqK80cfZaUQ1TotEu2eDoIPFockbUC"}, // password123
			{ID: "b5ed9407-681b-4dbb-b2d3-997803e8bbfc", Email: "jane@example.com", Password: "$2a$14$Od/6Z6WvfnaRAFPlzsaEEuSgOfStbdAnBO20vpQYhjnK1TNzmJHmS"}, // securepass
		},
	}
}

// GetByEmail retrieves a user by email
func (r *inMemoryRepository) GetByEmail(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, pkg.ErrUserNotFound
}

// GetByID retrieves a user by ID
func (r *inMemoryRepository) GetByID(id string) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, pkg.ErrUserNotFound
}

// GetAll retrieves all users (returns DTOs without passwords)
func (r *inMemoryRepository) GetAll() ([]UserDTO, error) {
	dtos := make([]UserDTO, 0, len(r.users))
	for _, user := range r.users {
		dtos = append(dtos, user.ToDTO())
	}
	return dtos, nil
}

// Create creates a new user with hashed password
func (r *inMemoryRepository) Create(email, password string) (*User, error) {
	// Check if user already exists
	if _, err := r.GetByEmail(email); err == nil {
		return nil, pkg.ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	// Create user
	user := User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: string(hashedPassword),
	}

	r.users = append(r.users, user)
	return &user, nil
}

// VerifyCredentials checks if email/password combination is valid
func (r *inMemoryRepository) VerifyCredentials(email, password string) (*UserDTO, error) {
	user, err := r.GetByEmail(email)
	if err != nil {
		return nil, pkg.ErrInvalidCredentials
	}

	// Compare password with hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, pkg.ErrInvalidCredentials
	}

	dto := user.ToDTO()
	return &dto, nil
}
