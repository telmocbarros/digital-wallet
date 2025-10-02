package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ACCESS_TOKEN_SECRET string
var REFRESH_TOKEN_SECRET string

func init() {
	// Find .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")
}
