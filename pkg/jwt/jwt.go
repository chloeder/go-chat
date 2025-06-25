package jwt

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
)

type ClaimToken struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

var matTypeToken = map[string]time.Duration{
	"token":   3 * time.Hour,
	"refresh": 72 * time.Hour,
}

func GenerateToken(ctx context.Context, username string, fullName string, typeToken string) (string, error) {
	secretKey := env.GetEnv("JWT_SECRET", "secret")

	claims := ClaimToken{
		Username: username,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(matTypeToken[typeToken])),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resultToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}

	return resultToken, nil
}
