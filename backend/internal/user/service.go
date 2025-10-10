package user

import (
	"digitalwallet/backend/pkg"
)

// Service handles user business logic
type Service struct {
	repo Repository
}

// NewService creates a new user service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Register creates a new user account
func (s *Service) Register(email, password, firstName, lastName string) (*pkg.UserDTO, error) {
	// Validate input
	if email == "" || password == "" || firstName == "" || lastName == "" {
		return nil, pkg.ErrMissingField
	}

	// Create user through repository
	user, err := s.repo.Create(email, password, firstName, lastName)
	if err != nil {
		return nil, err
	}

	dto := user.ToDTO()
	return &dto, nil
}

// Login authenticates a user with email and password
func (s *Service) Authenticate(email, password string) (*pkg.UserDTO, error) {
	// Validate input
	if email == "" || password == "" {
		return nil, pkg.ErrMissingField
	}

	// Verify credentials
	userDTO, err := s.repo.VerifyCredentials(email, password)
	if err != nil {
		return nil, err
	}

	return userDTO, nil
}

// GetByID retrieves a user by ID
func (s *Service) GetByID(id string) (*pkg.UserDTO, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	dto := user.ToDTO()
	return &dto, nil
}

// GetAll retrieves all users
func (s *Service) GetAll() ([]pkg.UserDTO, error) {
	return s.repo.GetAll()
}
