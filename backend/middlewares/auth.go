package middlewares

import (
	"digitalwallet/backend/config"
	"digitalwallet/backend/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	jwtAuthMiddleware(c)
	// sessionAuthMiddleware(c)
	c.Next()
}

func jwtAuthMiddleware(c *gin.Context) {
	// Retrieve the cookie from the request
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("Error retrieving cookie:", err)
		c.Abort()
		return
	}

	// Extract the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte(config.ACCESS_TOKEN_SECRET)
		return hmacSampleSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("Error parsing token:", err)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("Invalid token claims")
		c.Abort()
		return
	}

	// Check expiration and validity of the token
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		log.Println("Token expired", claims["exp"].(float64), time.Now().Unix())
		c.Abort()
		return
	}
	// Extract user information from the token
	userId := claims["userId"].(string)

	// Set user information in the context
	c.Set("userId", userId)
	// Go to the next middleware/handler
	c.Next()
}

func sessionAuthMiddleware(c *gin.Context) {
	// Retrieve the cookie from the request
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("Error retrieving cookie:", err)
		c.Abort()
		return
	}

	// Validate the session token
	session, exists := database.GetSession(tokenString)
	if !exists || session.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("Invalid or expired session token")
		c.Abort()
		return
	}

	// Fetch user from the database (mocked here)
	user, found := database.GetUserById(session.UserId)
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("User not found")
		c.Abort()
		return
	}

	// Set user information in the context
	c.Set("user", user)
	// Go to the next middleware/handler
	c.Next()
}
