package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	// Start server on port 8080 (default)
	fmt.Print("Server started at PORT 8080")
	r.Run()
}
