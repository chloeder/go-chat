package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
)

func CreateUser(ctx *fiber.Ctx, user *models.User) error {
	return database.DB.Create(user).Error
}

func CreateUserSession(ctx *fiber.Ctx, userSession *models.UserSession) error {
	return database.DB.Create(userSession).Error
}

func DeleteUserSessionByToken(ctx *fiber.Ctx, token string) error {
	return database.DB.Where("token = ?", token).Delete(&models.UserSession{}).Error
}

func FindUserSessionByToken(ctx *fiber.Ctx, token string) (*models.UserSession, error) {
	userSession := new(models.UserSession)

	err := database.DB.Where("token = ?", token).First(userSession).Error
	if err != nil {
		return nil, err
	}

	return userSession, nil
}

func FindUserByUsername(ctx *fiber.Ctx, username string) (*models.User, error) {
	user := new(models.User)

	err := database.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
