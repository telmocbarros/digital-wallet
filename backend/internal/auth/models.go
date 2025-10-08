package auth

// RefreshToken represents a refresh token stored in the system
type RefreshToken struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Revoked   bool   `json:"revoked"`
}

// Session represents a user session (currently unused, but kept for future use)
type Session struct {
	UserID    string
	CreatedAt int64
	ExpiresAt int64
}

// TokenPair represents access and refresh tokens
type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
