package jwt

import (
	"context"
	"fmt"
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

var jwtSecret = env.GetEnv("JWT_SECRET", "secret")

func GenerateToken(ctx context.Context, username string, fullName string, typeToken string) (string, error) {
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

	resultToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}

	return resultToken, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {
	var (
		claimToken *ClaimToken
		ok         bool
	)

	jwtToken, err := jwt.ParseWithClaims(token, &ClaimToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claimToken, ok = jwtToken.Claims.(*ClaimToken); ok && jwtToken.Valid {
		return claimToken, nil
	}

	return nil, fmt.Errorf("invalid token")
}
