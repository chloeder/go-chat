package controllers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/app/repositories"
	"github.com/kooroshh/fiber-boostrap/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	registerRequest := new(models.RegisterRequest)

	err := ctx.BodyParser(registerRequest)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err = registerRequest.Validate()
	if err != nil {
		log.Println("Validation error:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	userModel := models.User{
		Username: registerRequest.Username,
		FullName: registerRequest.FullName,
		Password: string(hashPassword),
	}

	err = repositories.CreateUser(ctx.Context(), &userModel)
	if err != nil {
		log.Println("Error creating user:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    userModel,
	})
}

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)

	err := ctx.BodyParser(loginRequest)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err = loginRequest.Validate()
	if err != nil {
		log.Println("Validation error:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
		})
	}

	userModel, err := repositories.FindUserByUsername(ctx.Context(), loginRequest.Username)
	if err != nil {
		log.Println("Error finding user:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error finding user",
		})
	}

	if userModel == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginRequest.Password))
	if err != nil {
		log.Println("Invalid password:", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	now := time.Now()
	token, err := jwt.GenerateToken(ctx.Context(), userModel.Username, userModel.FullName, "token", now)
	if err != nil {
		log.Println("Error generating token:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating token",
		})
	}

	refreshToken, err := jwt.GenerateToken(ctx.Context(), userModel.Username, userModel.FullName, "refresh_token", now)
	if err != nil {
		log.Println("Error generating refresh token:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating refresh token",
		})
	}

	userSession := models.UserSession{
		UserID:              userModel.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(time.Hour * 3),
		RefreshTokenExpired: now.Add(time.Hour * 24),
	}

	err = repositories.CreateUserSession(ctx.Context(), &userSession)
	if err != nil {
		log.Println("Error creating user session:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user session",
		})
	}

	loginResponse := models.LoginResponse{
		Username:     userModel.Username,
		FullName:     userModel.FullName,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    loginResponse,
	})
}

func Logout(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	err := repositories.DeleteUserSessionByToken(ctx.Context(), token)
	if err != nil {
		log.Println("Error deleting user session:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting user session",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func RefreshToken(ctx *fiber.Ctx) error {
	now := time.Now()
	username := ctx.Locals("username").(string)
	fullName := ctx.Locals("full_name").(string)
	refreshToken := ctx.Get("Authorization")

	token, err := jwt.GenerateToken(ctx.Context(), username, fullName, "token", now)
	if err != nil {
		log.Println("Error generating token:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating token",
		})
	}

	err = repositories.UpdateUserSessionToken(ctx.Context(), token, refreshToken)
	if err != nil {
		log.Println("Error updating user session token:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating user session token",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    token,
	})
}
