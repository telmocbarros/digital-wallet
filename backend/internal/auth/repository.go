package auth

import "digitalwallet/backend/pkg"

// Repository defines the interface for auth data access
type Repository interface {
	// Refresh token operations
	SaveRefreshToken(jti string, token RefreshToken) error
	GetRefreshToken(jti string) (*RefreshToken, error)
	RevokeRefreshToken(jti string) error

	// Session operations (for future use)
	SaveSession(sessionID string, session Session) error
	GetSession(sessionID string) (*Session, error)
	DeleteSession(sessionID string) error
}

// inMemoryRepository implements Repository using in-memory storage
type inMemoryRepository struct {
	refreshTokens map[string]RefreshToken
	sessions      map[string]Session
}

// NewRepository creates a new auth repository
func NewRepository() Repository {
	return &inMemoryRepository{
		refreshTokens: make(map[string]RefreshToken),
		sessions:      make(map[string]Session),
	}
}

// SaveRefreshToken stores a refresh token
func (r *inMemoryRepository) SaveRefreshToken(jti string, token RefreshToken) error {
	r.refreshTokens[jti] = token
	return nil
}

// GetRefreshToken retrieves a refresh token by JTI
func (r *inMemoryRepository) GetRefreshToken(jti string) (*RefreshToken, error) {
	token, exists := r.refreshTokens[jti]
	if !exists {
		return nil, pkg.ErrRefreshTokenNotFound
	}
	return &token, nil
}

// RevokeRefreshToken marks a refresh token as revoked
func (r *inMemoryRepository) RevokeRefreshToken(jti string) error {
	token, exists := r.refreshTokens[jti]
	if !exists {
		return pkg.ErrRefreshTokenNotFound
	}
	token.Revoked = true
	r.refreshTokens[jti] = token
	return nil
}

// SaveSession stores a session
func (r *inMemoryRepository) SaveSession(sessionID string, session Session) error {
	r.sessions[sessionID] = session
	return nil
}

// GetSession retrieves a session by ID
func (r *inMemoryRepository) GetSession(sessionID string) (*Session, error) {
	session, exists := r.sessions[sessionID]
	if !exists {
		return nil, pkg.ErrUnauthorized
	}
	return &session, nil
}

// DeleteSession removes a session
func (r *inMemoryRepository) DeleteSession(sessionID string) error {
	delete(r.sessions, sessionID)
	return nil
}
