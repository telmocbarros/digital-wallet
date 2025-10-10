package wallet

import (
	"digitalwallet/backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, walletHandler *Handler, authMiddleware *auth.Middleware) {
	// Protected routes
	router.POST("/wallets", authMiddleware.Authenticate, walletHandler.Create)
	router.GET("/wallets/:ID", authMiddleware.Authenticate, walletHandler.Get)
	router.POST("/wallets/:walletID/cards", authMiddleware.Authenticate, walletHandler.CreateCard)
	router.POST("/wallets/:walletID/cards/cardID", authMiddleware.Authenticate, walletHandler.RemoveCard)
}
