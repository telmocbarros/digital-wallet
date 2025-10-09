package user

import (
	"digitalwallet/backend/internal/auth"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all user-related routes
func RegisterRoutes(router *gin.Engine, userHandler *Handler, authMiddleware *auth.Middleware) {
	// Public routes
	router.POST("/users", userHandler.Create)

	// Protected routes
	router.GET("/users", authMiddleware.Authenticate, userHandler.List)
}
