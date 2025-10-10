package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ACCESS_TOKEN_SECRET string
var REFRESH_TOKEN_SECRET string

func init() {
	// Load .env file (optional in production where env vars are set by platform)
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")
}
