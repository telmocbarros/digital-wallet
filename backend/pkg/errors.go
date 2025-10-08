package pkg

import "errors"

// User errors
var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)

// Auth errors
var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrRefreshTokenRevoked = errors.New("refresh token revoked")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)

// Validation errors
var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrMissingField       = errors.New("missing required field")
)
