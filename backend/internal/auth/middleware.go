package auth

import (
	"digitalwallet/backend/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware handles authentication for protected routes
type Middleware struct {
	service *Service
}

// NewMiddleware creates a new auth middleware
func NewMiddleware(service *Service) *Middleware {
	return &Middleware{service: service}
}

// Authenticate validates the JWT token and sets user context
func (m *Middleware) Authenticate(c *gin.Context) {
	// Retrieve the access token from cookie
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		log.Println("Error retrieving access_token cookie:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Validate token
	userID, err := m.service.ValidateAccessToken(tokenString)
	if err != nil {
		if err == pkg.ErrTokenExpired {
			log.Println("Access token expired")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		} else {
			log.Println("Invalid access token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		c.Abort()
		return
	}

	// Set user ID in context
	c.Set("userId", userID)
	c.Next()
}
