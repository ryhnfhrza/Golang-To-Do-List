package util

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

var JWT_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaim struct {
    Username string `json:"username"`
    ID       string `json:"id"`
    jwt.RegisteredClaims
}

