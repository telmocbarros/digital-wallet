package ledger

import (
	"digitalwallet/backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, ledgerHandler *Handler, authMiddleware *auth.Middleware) {
	// Public routes (read-only queries)
	// These could be protected with auth middleware if needed
	ledger := router.Group("/api/ledger")
	{
		// Balance and statement queries
		ledger.GET("/balance/:accountId", authMiddleware.Authenticate, ledgerHandler.GetBalance)
		ledger.GET("/statement/:accountId", authMiddleware.Authenticate, ledgerHandler.GetStatement)
		ledger.GET("/transaction/:transactionId", authMiddleware.Authenticate, ledgerHandler.GetTransactionDetails)

		// Verification endpoints (admin/debugging)
		ledger.POST("/verify/account/:accountId", authMiddleware.Authenticate, ledgerHandler.VerifyAccountBalance)
		ledger.POST("/verify/transaction/:transactionId", authMiddleware.Authenticate, ledgerHandler.VerifyTransaction)
	}
}
