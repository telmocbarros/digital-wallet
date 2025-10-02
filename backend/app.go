package main

import (
	"crypto/rand"
	"digitalwallet/backend/config"
	"digitalwallet/backend/database"
	"digitalwallet/backend/middlewares"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User = database.User

func main() {
	accessToken := config.ACCESS_TOKEN_SECRET

	if accessToken == "" {
		fmt.Println("No ACCESS_TOKEN_SECRET found")
		return
	}

	refreshToken := config.REFRESH_TOKEN_SECRET
	if refreshToken == "" {
		fmt.Println("No REFRESH_TOKEN_SECRET found")
		return
	}

	// Initialise the Gin router
	r := gin.Default()

	// Middleware to handle CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173") // Adjust this to your frontend's origin
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true") // This is crucial!
		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// GET endpoit at "/"
	r.GET("/", middlewares.AuthMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome back!"})
	})

	// POST endpoint to handle login
	r.POST("/login", func(c *gin.Context) {
		var email = c.Query("email")
		var password = c.Query("password")

		if email == "" || password == "" {
			fmt.Println("Invalid login attempt: missing email or password")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}

		//Check if the user exists in the users slice.
		userData, err := loginUser(email, password)
		if err {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		//Generate a JWT token for the user.
		// Currently using this one token both as identity and access token.
		accessToken := generateAccessToken(accessToken, userData)
		refreshToken := generateRefreshToken(refreshToken, userData)

		// token := generateSessionToken(userData)
		setCookie(c, "access_token", accessToken, 300)    // Cookie expires in 300 seconds (5 minutes)
		setCookie(c, "refresh_token", refreshToken, 6400) // Cookie expires in 6400 seconds (a bit less than 2 hours)
		//Add a 201 status code to the response, along with JSON representing the user that logged in.
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	// GET endpoint to fetch users
	r.GET("/users", middlewares.AuthMiddleware, func(c *gin.Context) {
		users := database.GetUsers()
		c.IndentedJSON(http.StatusOK, users)
	})

	r.POST("/users", middlewares.AuthMiddleware, func(c *gin.Context) {
		var parsedData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if c.BindJSON(&parsedData) != nil {
			log.Fatal("Invalid request: missing email or password (email: ", parsedData.Email, ", password: ", parsedData.Password, ")")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		if parsedData.Email == "" || parsedData.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}

		database.SaveUser(parsedData.Email, parsedData.Password)
	})

	// Start server on port 8080 (default)
	fmt.Print("Server started at PORT 8080")
	r.Run()
}

func loginUser(email, password string) (data database.UserDTO, error bool) {
	var user, isVerified = database.VerifyUserCredentials(email, password)
	if !isVerified {
		return database.UserDTO{}, true
	}
	return user, false
}

func setCookie(c *gin.Context, cookieName, token string, expiryTime int) {
	// Set the JWT token in a cookie
	// Cookie will expire in 5 minutes
	// In a real application, consider setting the Secure and SameSite attributes appropriately.
	c.SetCookie(cookieName, token, expiryTime, "/", "", false, true)
}

func generateAccessToken(accessToken string, user database.UserDTO) string {

	key := []byte(accessToken)
	// Create the Claims
	// Set token to expire in 5 minutes.
	// This is just for demonstration purposes. In a real application, you might want to set a shorter expiration time.
	// Also, consider using refresh tokens for better security.
	// See: https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-best-practices
	// and https://developer.okta.com/blog/2019/06/18/refresh-tokens-what-are-they-and-when-to-use-them
	// for more information.
	// Note: The "ttl" claim is not a standard JWT claim. It's used here for demonstration purposes only.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userEmail": user.Email,
		"userId":    user.Id,
		"exp":       time.Now().Add(time.Minute * 5).Unix(), // This represents 5 minutes from now
	})
	s, err := token.SignedString(key)

	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}
	return s
}

func generateSessionToken(user User) string {
	// Create a new session for the user
	sessionId := rand.Text()

	database.SaveSession(sessionId, database.Session{
		UserId:    user.Id,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Session expires in 24 hours
	})
	return sessionId
}

func generateRefreshToken(refreshToken string, user database.UserDTO) string {
	key := []byte(refreshToken)
	jti := rand.Text()
	expirationTime := time.Now().Add(7 * 24 * time.Hour).Unix() // Refresh token valid for 7 days
	creationTime := time.Now().Unix()
	// Create a new JWT token for the refresh token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"exp":    expirationTime,
		"jti":    jti, // Unique identifier for the token
	})

	refreshToken, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Error generating refresh token:", err)
		return ""
	}

	database.SaveRefreshToken(jti, database.RefreshToken{
		ID:        jti,
		Token:     refreshToken,
		ExpiresAt: expirationTime,
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
		Revoked:   false,
	})

	return refreshToken

}
