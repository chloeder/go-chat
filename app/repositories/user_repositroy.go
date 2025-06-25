package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
)

func CreateUser(ctx *fiber.Ctx, user *models.User) error {
	return database.DB.Create(user).Error
}
