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

// Login handles user login
// POST /login
func (h *Handler) Login(c *gin.Context) {
	var req pkg.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println("Invalid login request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Authenticate user
	userDTO, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		if err == pkg.ErrInvalidCredentials || err == pkg.ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Return user data (tokens will be set by auth handler)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    userDTO,
	})
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
