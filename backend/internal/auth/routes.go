package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all auth-related routes
func RegisterRoutes(router *gin.Engine, authHandler *Handler, authMiddleware *Middleware) {
	// Public routes
	router.GET("/auth/status", authHandler.Status)
	router.POST("/refresh", authHandler.Refresh)

	// Protected routes
	router.POST("/logout", authMiddleware.Authenticate, authHandler.Logout)

	// Protected test route
	router.GET("/", authMiddleware.Authenticate, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome back!"})
	})
}
