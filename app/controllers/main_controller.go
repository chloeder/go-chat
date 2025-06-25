package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

func RenderHello(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	err := ctx.BodyParser(user)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err = user.Validate()
	if err != nil {
		log.Println("Validation error:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	user.Password = string(hashPassword)

	err = repositories.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error creating user:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}
