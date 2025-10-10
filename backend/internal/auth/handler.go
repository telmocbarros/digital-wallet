package auth

import (
	"digitalwallet/backend/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for auth operations
type Handler struct {
	service *Service
}

// NewHandler creates a new auth handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Status checks if the user is authenticated
// GET /auth/status
func (h *Handler) Status(c *gin.Context) {
	// Try to get the access token cookie
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authenticated": false})
		return
	}

	// Validate token
	userID, err := h.service.ValidateAccessToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authenticated": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"userId":        userID,
	})
}

// Refresh generates new access and refresh tokens
// POST /refresh
func (h *Handler) Refresh(c *gin.Context) {
	// Get refresh token from cookie
	refreshTokenCookie, err := c.Cookie("refresh_token")
	if err != nil {
		log.Println("Error retrieving refresh token cookie:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Refresh tokens
	tokenPair, err := h.service.RefreshTokens(refreshTokenCookie)
	if err != nil {
		log.Println("Error refreshing tokens:", err)
		// Clear cookies on refresh failure
		h.clearTokenCookies(c)

		if err == pkg.ErrTokenExpired {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		} else if err == pkg.ErrRefreshTokenRevoked {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token revoked"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		return
	}

	// Set new tokens in cookies
	h.setTokenCookies(c, tokenPair.AccessToken, tokenPair.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed successfully"})
}

// Login
// POST /login
func (h *Handler) Login(c *gin.Context) {
	var req pkg.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Authenticate user
	userID, userEmail, err := h.service.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate tokens
	tokenPair, err := h.service.GenerateTokens(userID, userEmail)
	if err != nil {
		log.Println("Error generating tokens:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Set tokens as cookies
	h.SetTokenCookies(c, tokenPair.AccessToken, tokenPair.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Logout revokes the refresh token and clears cookies
// POST /logout (requires authentication)
func (h *Handler) Logout(c *gin.Context) {
	userID := c.GetString("userId")
	log.Printf("User %s logged out successfully", userID)

	// Get refresh token to revoke it
	refreshTokenCookie, err := c.Cookie("refresh_token")
	if err != nil {
		log.Println("No refresh token found, clearing cookies only")
		h.clearTokenCookies(c)
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
		return
	}

	// Revoke refresh token
	if err := h.service.RevokeRefreshTokenByString(refreshTokenCookie); err != nil {
		log.Println("Error revoking refresh token:", err)
	}

	// Clear cookies
	h.clearTokenCookies(c)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// SetTokenCookies is a helper to set auth tokens as cookies (used by other handlers)
func (h *Handler) SetTokenCookies(c *gin.Context, accessToken, refreshToken string) {
	h.setTokenCookies(c, accessToken, refreshToken)
}

// setTokenCookies sets access and refresh token cookies
func (h *Handler) setTokenCookies(c *gin.Context, accessToken, refreshToken string) {
	c.SetCookie("access_token", accessToken, 30, "/", "", false, true)    // 30 seconds
	c.SetCookie("refresh_token", refreshToken, 300, "/", "", false, true) // 300 seconds (5 minutes)
}

// clearTokenCookies removes access and refresh token cookies
func (h *Handler) clearTokenCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}
