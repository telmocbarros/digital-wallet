package main

import (
	"digitalwallet/backend/config"
	"digitalwallet/backend/database"
	"digitalwallet/backend/middlewares"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User = database.User

var users = database.Data

func main() {

	jwtSecret := config.JWT_SECRET
	if jwtSecret == "" {
		fmt.Println("No JWT_SECRET found")
		return
	}
	// Initialise the Gin router
	r := gin.Default()

	// Middleware to handle CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	// GET endpoit at "/"
	r.GET("/", middlewares.AuthMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// POST endpoint to handle login
	r.POST("/login", func(c *gin.Context) {
		var user User

		//Use Context.BindJSON to bind the request body to user.
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		//Check if the user exists in the users slice.
		result, err := loginUser(user)
		if err {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		//Generate a JWT token for the user.
		token := generateToken(jwtSecret, result)
		setCookie(c, token)

		//Add a 201 status code to the response, along with JSON representing the user that logged in.
		c.IndentedJSON(http.StatusOK, user)
	})

	// GET endpoint to fetch users
	r.GET("/users", middlewares.AuthMiddleware, func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})

	// Start server on port 8080 (default)
	fmt.Print("Server started at PORT 8080")
	r.Run()
}

func loginUser(user User) (data User, error bool) {
	for _, value := range users {
		if value.Email == user.Email && value.Password == user.Password {
			return value, false
		}
	}
	return User{}, true
}

func generateToken(jwtSecret string, user User) string {

	var (
		key []byte
		t   *jwt.Token
	)

	key = []byte(jwtSecret)
	// Create the Claims
	// Set token to expire in 100 days
	// This is just for demonstration purposes. In a real application, you might want to set a shorter expiration time.
	// Also, consider using refresh tokens for better security.
	// See: https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-best-practices
	// and https://developer.okta.com/blog/2019/06/18/refresh-tokens-what-are-they-and-when-to-use-them
	// for more information.
	// Note: The "ttl" claim is not a standard JWT claim. It's used here for demonstration purposes only.
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userEmail": user.Email,
		"userId":    user.Id,
		"ttl":       time.Now().Add(time.Hour * 24 * 100).Unix(), // This represents 100 days in the future
	})
	s, err := t.SignedString(key)

	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}
	return s
}

func setCookie(c *gin.Context, token string) {
	c.SetCookie("Authorization", token, 3600*24*100, "", "", false, true)
}

func generateSession(user User) string {
	return ""
}
