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
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ACCESS_TOKEN_EXPIRY = time.Second * 30 // 30 seconds for testing purposes
var REFRESH_TOKEN_EXPIRY = time.Minute * 5 // 1 minute for testing purposes

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

	// GET endpoint to check authentication status
	r.GET("/auth/status", func(c *gin.Context) {
		// Try to get the access token cookie
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			// No token present - user is not authenticated
			c.JSON(http.StatusOK, gin.H{"authenticated": false})
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.ACCESS_TOKEN_SECRET), nil
		})

		if err != nil || !token.Valid {
			// Invalid token
			c.JSON(http.StatusOK, gin.H{"authenticated": false})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"authenticated": false})
			return
		}

		// Check expiration
		if int64(claims["exp"].(float64)) < time.Now().Unix() {
			c.JSON(http.StatusOK, gin.H{"authenticated": false})
			return
		}

		// Token is valid
		userId := claims["userId"].(string)
		c.JSON(http.StatusOK, gin.H{
			"authenticated": true,
			"userId":        userId,
		})
	})

	// POST endpoint to handle login
	r.POST("/login", func(c *gin.Context) {
		//Parse the JSON body to get email and password.
		var user UserLogin
		if err := c.BindJSON(&user); err != nil {
			fmt.Println("Invalid login attempt: missing email or password", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}

		email := user.Email
		password := user.Password

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
		setCookie(c, "access_token", accessToken, 30)    // Cookie expires in 300 seconds (5 minutes)
		setCookie(c, "refresh_token", refreshToken, 300) // Cookie expires in 6400 seconds (a bit less than 2 hours)
		//Add a 201 status code to the response, along with JSON representing the user that logged in.
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	r.POST("/logout", middlewares.AuthMiddleware, func(c *gin.Context) {
		userId := c.GetString("userId")

		// Here you can implement any server-side session invalidation if needed.
		// For example, you might want to store a blacklist of tokens or sessions.

		log.Printf("User %s logged out successfully", userId)

		// Get refresh token to revoke it
		refreshTokenCookie, err := c.Cookie("refresh_token")
		if err != nil {
			log.Println("No refresh token found, clearing cookies only")
			setCookie(c, "access_token", "", -1)
			setCookie(c, "refresh_token", "", -1)
			c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
			return
		}

		// Parse refresh token to get JTI
		rToken, err := jwt.Parse(refreshTokenCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.REFRESH_TOKEN_SECRET), nil
		})

		if err == nil {
			if rtClaims, ok := rToken.Claims.(jwt.MapClaims); ok {
				jti := rtClaims["jti"].(string)
				database.RevokeRefreshToken(jti)
			}
		}

		// Clear the access token cookie
		setCookie(c, "access_token", "", -1)  // Set expiry to -1 to delete the cookie
		setCookie(c, "refresh_token", "", -1) // Set expiry to -1 to delete the cookie

		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	})

	// POST endpoint to handle refresh token
	r.POST("/refresh", func(c *gin.Context) {
		refreshTokenCookie, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			log.Println("Error retrieving refresh token cookie:", err)
			return
		}

		// Parse and validate the refresh token
		var accessToken, refreshToken = handleRefreshToken(c, refreshTokenCookie, refreshToken)
		if accessToken == "" || refreshToken == "" {
			// Clear cookies when refresh fails
			setCookie(c, "access_token", "", -1)
			setCookie(c, "refresh_token", "", -1)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Set the new access token in the cookie
		setCookie(c, "access_token", accessToken, 300)    // Cookie expires in 300 seconds (5 minutes)
		setCookie(c, "refresh_token", refreshToken, 6400) // Cookie expires in 6400 seconds (a bit less than 2 hours)

		// Respond with a success message
		c.JSON(http.StatusOK, gin.H{"message": "Access token refreshed successfully"})

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

func generateAccessToken(accessTokenSecret string, user database.UserDTO) string {

	key := []byte(accessTokenSecret)
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
		"exp":       time.Now().Add(ACCESS_TOKEN_EXPIRY).Unix(), // This represents 30 seconds from now
		//"exp": time.Now().Add(time.Minute * 5).Unix(), // This represents 5 minutes from now
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
		ExpiresAt: time.Now().Add(ACCESS_TOKEN_EXPIRY).Unix(), // This represents 30 seconds from now,
		//ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Session expires in 24 hours
	})
	return sessionId
}

func generateRefreshToken(refreshTokenSecret string, user database.UserDTO) string {
	key := []byte(refreshTokenSecret)
	jti := rand.Text()
	//expirationTime := time.Now().Add(7 * 24 * time.Hour).Unix()   // Refresh token valid for 7 days
	expirationTime := time.Now().Add(REFRESH_TOKEN_EXPIRY).Unix() // Refresh token valid for 1 minute
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

func handleRefreshToken(c *gin.Context, refreshTokenCookie, refreshTokenSecret string) (string, string) {
	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshTokenCookie, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshTokenSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Println("Error parsing refresh token:", err)
		return "", ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.Println("Invalid refresh token claims")
		return "", ""
	}

	// Check expiration of the refresh token
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		log.Println("Refresh token expired")
		return "", ""
	}

	// Check if the refresh token is revoked
	jti := claims["jti"].(string)
	savedToken, exists := database.GetRefreshToken(jti)
	if !exists || savedToken.Revoked {
		log.Println("Refresh token revoked or not found")
		return "", ""
	}

	// Generate a new access token
	var newAccessToken = generateAccessToken(refreshTokenSecret, database.UserDTO{Id: claims["userId"].(string)})
	if newAccessToken == "" {
		log.Println("Error generating new access token")
		return "", ""
	}

	// Revoke the used refresh token
	database.RevokeRefreshToken(jti)

	// Generate a new refresh token
	var newRefreshToken = generateRefreshToken(refreshTokenSecret, database.UserDTO{Id: claims["userId"].(string)})
	if newRefreshToken == "" {
		log.Fatal("Error generating new refresh token")
		return "", ""
	}

	return newAccessToken, newRefreshToken
}
