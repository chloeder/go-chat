package repositories

import (
	"context"

	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
)

func CreateUser(ctx context.Context, user *models.User) error {
	return database.DB.Create(user).Error
}

func CreateUserSession(ctx context.Context, userSession *models.UserSession) error {
	return database.DB.Create(userSession).Error
}

func DeleteUserSessionByToken(ctx context.Context, token string) error {
	return database.DB.Where("token = ?", token).Delete(&models.UserSession{}).Error
}

func UpdateUserSessionToken(ctx context.Context, token string, refreshToken string) error {
	return database.DB.Exec("UPDATE user_sessions SET token = ? WHERE refresh_token = ?", token, refreshToken).Error
}

func FindUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	userSession := new(models.UserSession)

	err := database.DB.Where("token = ?", token).First(userSession).Error
	if err != nil {
		return nil, err
	}

	return userSession, nil
}

func FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := new(models.User)

	err := database.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
