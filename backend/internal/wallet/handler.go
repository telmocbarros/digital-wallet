package wallet

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

func (h *Handler) Create(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Println("User is not set in the user context:", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not authenticated"})
		return
	}

	walletID, err := h.service.CreateWallet(userId)
	if err != nil {
		log.Println("Error creating a wallet:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Wallet created successfully", "wallet_id": walletID})

}

func (h *Handler) Get(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Println("User is not set in the user context:", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not authenticated"})
		return
	}

	walletID := c.Param("walletID")
	if walletID == "" {
		log.Println("Error: Request missing required parameter (ID)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	wallet, err := h.service.GetWalletByID(walletID)
	if err != nil {
		log.Println("Error creating a wallet:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"wallet": wallet,
	})

}

func (h *Handler) CreateCard(c *gin.Context) {
	userId := c.GetString("userId")

	if userId == "" {
		log.Println("Error: User is not set in the user context", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not authenticated"})
		return
	}
	walletID := c.Param("walletID")
	if walletID == "" {
		log.Println("Error: Request missing required parameter (walletID)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	var card *CardDTO
	if err := c.BindJSON(&card); err != nil {
		log.Println("Error: binding the request payload to the CardDTO struct:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	cardId, err := h.service.AddCard(walletID, card)
	if err != nil {
		log.Println("Error: adding a card to the wallet:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Card added successfully", "card_id": cardId})

}

func (h *Handler) GetCard(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Println("Error: User is not set in the user context", userId)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	walletID, cardID := c.Param("walletID"), c.Param("cardID")
	if cardID == "" || walletID == "" {
		log.Println("Error: Request missing required parameters (walletID, cardID)")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	card, err := h.service.GetCard(walletID, cardID)
	if err != nil {
		log.Println("Error: getting card details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"card": card,
	})
}

func (h *Handler) RemoveCard(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Println()
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	walletID, cardID := c.Param("walletID"), c.Param("cardID")
	if cardID == "" || walletID == "" {
		log.Println()
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Requested": "Required parameters are not set"})
		return
	}

	if err := h.service.RemoveCard(walletID, cardID); err != nil {
		log.Println()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internaval Server Error"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
