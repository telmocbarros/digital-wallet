package auth

import (
	"crypto/rand"
	"digitalwallet/backend/pkg"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpiry  = 30 * time.Second // 30 seconds for testing
	RefreshTokenExpiry = 5 * time.Minute  // 5 minutes for testing
)

// Service handles authentication business logic
type Service struct {
	repo                 Repository
	accessTokenSecret    string
	refreshTokenSecret   string
}

// NewService creates a new auth service
func NewService(repo Repository, accessTokenSecret, refreshTokenSecret string) *Service {
	return &Service{
		repo:                 repo,
		accessTokenSecret:    accessTokenSecret,
		refreshTokenSecret:   refreshTokenSecret,
	}
}

// GenerateTokens creates a new access and refresh token pair for a user
func (s *Service) GenerateTokens(userID, userEmail string) (*TokenPair, error) {
	accessToken, err := s.generateAccessToken(userID, userEmail)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// generateAccessToken creates a new JWT access token
func (s *Service) generateAccessToken(userID, userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"userEmail": userEmail,
		"exp":       time.Now().Add(AccessTokenExpiry).Unix(),
	})

	return token.SignedString([]byte(s.accessTokenSecret))
}

// generateRefreshToken creates a new JWT refresh token and stores it
func (s *Service) generateRefreshToken(userID string) (string, error) {
	jti := generateRandomString()
	expirationTime := time.Now().Add(RefreshTokenExpiry).Unix()
	creationTime := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"exp":    expirationTime,
		"jti":    jti,
	})

	refreshToken, err := token.SignedString([]byte(s.refreshTokenSecret))
	if err != nil {
		return "", err
	}

	// Store refresh token
	err = s.repo.SaveRefreshToken(jti, RefreshToken{
		ID:        jti,
		UserID:    userID,
		Token:     refreshToken,
		ExpiresAt: expirationTime,
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
		Revoked:   false,
	})

	return refreshToken, err
}

// ValidateAccessToken validates an access token and returns the user ID
func (s *Service) ValidateAccessToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.accessTokenSecret), nil
	})

	if err != nil {
		return "", pkg.ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", pkg.ErrTokenInvalid
	}

	// Check expiration
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return "", pkg.ErrTokenExpired
	}

	userID := claims["userId"].(string)
	return userID, nil
}

// RefreshTokens validates a refresh token and generates new token pair
func (s *Service) RefreshTokens(refreshTokenString string) (*TokenPair, error) {
	// Parse refresh token
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.refreshTokenSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, pkg.ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, pkg.ErrTokenInvalid
	}

	// Check expiration
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return nil, pkg.ErrTokenExpired
	}

	// Check if token is revoked
	jti := claims["jti"].(string)
	savedToken, err := s.repo.GetRefreshToken(jti)
	if err != nil {
		return nil, pkg.ErrRefreshTokenNotFound
	}

	if savedToken.Revoked {
		return nil, pkg.ErrRefreshTokenRevoked
	}

	// Revoke old refresh token
	s.repo.RevokeRefreshToken(jti)

	// Generate new token pair
	userID := claims["userId"].(string)
	return s.GenerateTokens(userID, "")
}

// RevokeRefreshTokenByString parses a refresh token and revokes it
func (s *Service) RevokeRefreshTokenByString(refreshTokenString string) error {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.refreshTokenSecret), nil
	})

	if err != nil {
		// Token is invalid, but we can still try to revoke it
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		jti := claims["jti"].(string)
		return s.repo.RevokeRefreshToken(jti)
	}

	return nil
}

// generateRandomString creates a random string for JTI
func generateRandomString() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
