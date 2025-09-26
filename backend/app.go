package main

import (
	"digitalwallet/backend/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User = database.User

var users = database.Data

func main() {
	// Initialise the Gin router
	r := gin.Default()

	// Middleware to handle CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	// GET endpoit at "/"
	r.GET("/", func(c *gin.Context) {
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
		if !loginUser(user) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		//Add a 201 status code to the response, along with JSON representing the user that logged in.
		c.IndentedJSON(http.StatusOK, user)
	})

	// GET endpoint to fetch users
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})

	// Start server on port 8080 (default)
	fmt.Print("Server started at PORT 8080")
	r.Run()
}

func loginUser(user User) bool {
	for _, value := range users {
		if value.Email == user.Email && value.Password == user.Password {
			return true
		}
	}
	return false
}
