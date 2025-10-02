package database

type RefreshToken struct {
	ID string `json:"id"`
	// UserID is the ID of the user associated with the refresh token
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	// ExpiresAt is the expiration time of the refresh token (Unix timestamp)
	ExpiresAt int64 `json:"expires_at"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Revoked   bool  `json:"revoked"`
}

var refreshTokenStorage = make(map[string]RefreshToken)

func SaveRefreshToken(tokenId string, refreshToken RefreshToken) {
	refreshTokenStorage[tokenId] = refreshToken
}

func GetRefreshToken(tokenId string) (RefreshToken, bool) {
	refreshToken, exists := refreshTokenStorage[tokenId]
	return refreshToken, exists
}

func RevokeRefreshToken(tokenId string) {
	if token, exists := refreshTokenStorage[tokenId]; exists {
		token.Revoked = true
		refreshTokenStorage[tokenId] = token
	}
}
