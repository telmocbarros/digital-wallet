package main

import (
	"digitalwallet/backend/config"
	"digitalwallet/backend/internal/auth"
	"digitalwallet/backend/internal/ledger"
	"digitalwallet/backend/internal/user"
	"digitalwallet/backend/internal/wallet"
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
	walletRepo := wallet.NewRepository()
	ledgerRepo := ledger.NewRepository()

	// Initialize services
	userService := user.NewService(userRepo)
	authService := auth.NewService(authRepo, userService, config.ACCESS_TOKEN_SECRET, config.REFRESH_TOKEN_SECRET)
	walletService := wallet.NewService(walletRepo)
	ledgerService := ledger.NewService(ledgerRepo)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	authMiddleware := auth.NewMiddleware(authService)
	userHandler := user.NewHandler(userService)
	walletHandler := wallet.NewHandler(walletService)
	ledgerHandler := ledger.NewHandler(ledgerService)

	// Register routes
	auth.RegisterRoutes(r, authHandler, authMiddleware)
	user.RegisterRoutes(r, userHandler, authMiddleware)
	wallet.RegisterRoutes(r, walletHandler, authMiddleware)
	ledger.RegisterRoutes(r, ledgerHandler, authMiddleware)

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
