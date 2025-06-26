package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/repositories"
	"github.com/kooroshh/fiber-boostrap/pkg/jwt"
)

func MiddlewareValidateAuth(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")

	if auth == "" {
		log.Println("Unauthorized: No token provided")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	_, err := repositories.FindUserSessionByToken(ctx, auth)
	if err != nil {
		log.Println("Unauthorized: User session not found")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claimToken, err := jwt.ValidateToken(ctx.Context(), auth)
	if err != nil {
		log.Println("Unauthorized: Invalid token", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if time.Now().Unix() > claimToken.ExpiresAt.Unix() {
		log.Println("Unauthorized: Token expired")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	ctx.Set("username", claimToken.Username)
	ctx.Set("full_name", claimToken.FullName)

	return ctx.Next()
}