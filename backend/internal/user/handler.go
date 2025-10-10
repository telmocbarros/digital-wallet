package user

import (
	"digitalwallet/backend/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for user operations
type Handler struct {
	service *Service
}

// NewHandler creates a new user handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Create handles user registration
// POST /users
func (h *Handler) Create(c *gin.Context) {
	var req pkg.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println("Invalid create user request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, password, first name, and last name are required"})
		return
	}

	// Register user
	userDTO, err := h.service.Register(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		if err == pkg.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    userDTO,
	})
}

// List retrieves all users
// GET /users
func (h *Handler) List(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		log.Println("Error retrieving users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Retrieve a single user
// GET /users/id
func (h *Handler) GetById(c *gin.Context) {
	requestUserId := c.Query("userId")
	ctxUserId := c.GetString("userId")

	if ctxUserId != requestUserId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	userDTO, err := h.service.GetByID(requestUserId)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Println("Error retrieving user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, userDTO)
}
