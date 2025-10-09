package main

import (
	"digitalwallet/backend/config"
	"digitalwallet/backend/internal/auth"
	"digitalwallet/backend/internal/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Validate configuration
	if config.ACCESS_TOKEN_SECRET == "" {
		log.Fatal("No ACCESS_TOKEN_SECRET found")
	}
	if config.REFRESH_TOKEN_SECRET == "" {
		log.Fatal("No REFRESH_TOKEN_SECRET found")
	}

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(corsMiddleware())

	// Initialize repositories
	userRepo := user.NewRepository()
	authRepo := auth.NewRepository()

	// Initialize services
	userService := user.NewService(userRepo)
	authService := auth.NewService(authRepo, config.ACCESS_TOKEN_SECRET, config.REFRESH_TOKEN_SECRET)

	// Initialize handlers
	userHandler := user.NewHandler(userService)
	authHandler := auth.NewHandler(authService)
	authMiddleware := auth.NewMiddleware(authService)

	// Register routes
	auth.RegisterRoutes(r, authHandler, authMiddleware)
	user.RegisterRoutes(r, userHandler, authMiddleware)

	// Login route (combines user auth + token generation)
	r.POST("/login", func(c *gin.Context) {
		var req user.LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}

		// Authenticate user
		userDTO, err := userService.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate tokens
		tokenPair, err := authService.GenerateTokens(userDTO.ID, userDTO.Email)
		if err != nil {
			log.Println("Error generating tokens:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Set tokens as cookies
		authHandler.SetTokenCookies(c, tokenPair.AccessToken, tokenPair.RefreshToken)

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	// Start server
	fmt.Println("Server started at PORT 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// corsMiddleware handles CORS for the frontend
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
