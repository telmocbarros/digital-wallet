package ledger

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetBalance retrieves the current balance for an account
// GET /api/ledger/balance/:accountId
func (h *Handler) GetBalance(c *gin.Context) {
	accountID := c.Param("accountId")
	if accountID == "" {
		log.Println("Error: Missing required parameter (accountId)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter: accountId"})
		return
	}

	balance, err := h.service.GetBalance(accountID)
	if err != nil {
		if err == ErrAccountBalanceNotFound {
			log.Printf("Account balance not found: %s", accountID)
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}
		log.Printf("Error getting balance for account %s: %v", accountID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": balance.AccountID,
		"balance":    balance.ToDTO(),
	})
}

// GetStatement retrieves all ledger entries for an account (account statement)
// GET /api/ledger/statement/:accountId
func (h *Handler) GetStatement(c *gin.Context) {
	accountID := c.Param("accountId")
	if accountID == "" {
		log.Println("Error: Missing required parameter (accountId)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter: accountId"})
		return
	}

	entries, err := h.service.GetAccountStatement(accountID)
	if err != nil {
		log.Printf("Error getting statement for account %s: %v", accountID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Convert entries to DTOs for user-friendly response
	entryDTOs := make([]*LedgerEntryDTO, len(entries))
	for i, entry := range entries {
		entryDTOs[i] = entry.ToDTO()
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": accountID,
		"entries":    entryDTOs,
		"count":      len(entryDTOs),
	})
}

// GetTransactionDetails retrieves all ledger entries for a specific transaction
// GET /api/ledger/transaction/:transactionId
func (h *Handler) GetTransactionDetails(c *gin.Context) {
	transactionID := c.Param("transactionId")
	if transactionID == "" {
		log.Println("Error: Missing required parameter (transactionId)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter: transactionId"})
		return
	}

	entries, err := h.service.GetTransactionDetails(transactionID)
	if err != nil {
		log.Printf("Error getting transaction details for %s: %v", transactionID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if len(entries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Convert entries to DTOs
	entryDTOs := make([]*LedgerEntryDTO, len(entries))
	for i, entry := range entries {
		entryDTOs[i] = entry.ToDTO()
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transactionID,
		"entries":        entryDTOs,
		"count":          len(entryDTOs),
	})
}

// VerifyAccountBalance verifies that cached balance matches calculated balance
// POST /api/ledger/verify/account/:accountId
// This is useful for admin/debugging purposes
func (h *Handler) VerifyAccountBalance(c *gin.Context) {
	accountID := c.Param("accountId")
	if accountID == "" {
		log.Println("Error: Missing required parameter (accountId)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter: accountId"})
		return
	}

	balanced, err := h.service.VerifyAccountBalance(accountID)
	if err != nil {
		log.Printf("Error verifying account %s: %v", accountID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if !balanced {
		log.Printf("WARNING: Account %s has mismatched balances!", accountID)
		c.JSON(http.StatusOK, gin.H{
			"account_id": accountID,
			"verified":   false,
			"message":    "Cached balance does not match calculated balance",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": accountID,
		"verified":   true,
		"message":    "Account balance verified successfully",
	})
}

// VerifyTransaction verifies that all entries for a transaction sum to zero
// POST /api/ledger/verify/transaction/:transactionId
// This is useful for admin/debugging purposes
func (h *Handler) VerifyTransaction(c *gin.Context) {
	transactionID := c.Param("transactionId")
	if transactionID == "" {
		log.Println("Error: Missing required parameter (transactionId)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter: transactionId"})
		return
	}

	err := h.service.VerifyTransaction(transactionID)
	if err != nil {
		if err == ErrTransactionNotBalanced {
			log.Printf("WARNING: Transaction %s does not balance!", transactionID)
			c.JSON(http.StatusOK, gin.H{
				"transaction_id": transactionID,
				"verified":       false,
				"message":        "Transaction entries do not sum to zero",
			})
			return
		}
		log.Printf("Error verifying transaction %s: %v", transactionID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transactionID,
		"verified":       true,
		"message":        "Transaction verified successfully",
	})
}
