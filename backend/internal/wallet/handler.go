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

	if err := h.service.CreateWallet(userId); err != nil {
		log.Println("Error creating a wallet:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, nil)

}

func (h *Handler) Get(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Println("User is not set in the user context:", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not authenticated"})
		return
	}

	walletID := c.Query("ID")
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
	// userId := c.GetString("userId")

	// if userId == "" {
	// 	log.Println("User is not set in the user context:", userId)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "User not authenticated"})
	// 	return
	// }

	// if err := h.service.CreateWallet(userId); err != nil {
	// 	log.Println("Error creating a wallet:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	// 	return
	// }
	c.JSON(http.StatusOK, nil)

}

func (h *Handler) RemoveCard(c *gin.Context) {
	//userId := c.GetString("userId")
	c.JSON(http.StatusOK, nil)
}
