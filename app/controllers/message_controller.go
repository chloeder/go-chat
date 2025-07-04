package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/repositories"
)

func GetMessages(ctx *fiber.Ctx) error {
	messages, err := repositories.GetMessages(ctx.Context())
	if err != nil {
		log.Println("Error getting messages:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    messages,
	})
}
