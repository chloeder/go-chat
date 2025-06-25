package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	Username  string    `gorm:"unique;" validate:"required,min=6,max=32"`
	FullName  string    `gorm:"type:text;" validate:"required,min=6,max=32"`
	Password  string    `gorm:"type:text;" validate:"required,min=6"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserSession struct {
	ID                  uint      `gorm:"primarykey"`
	UserID              uint      `json:"user_id" gorm:"type:int;not null" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(255);not null" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(255);not null" validate:"required"`
	TokenExpired        time.Time `gorm:"type:datetime;not null" validate:"required"`
	RefreshTokenExpired time.Time `gorm:"type:datetime;not null" validate:"required"`
	CreatedAt           time.Time `gorm:"autoCreateTime"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime"`
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=6,max=32"`
		FullName string `json:"full_name" validate:"required,min=6,max=32"`
		Password string `json:"password" validate:"required,min=6"`
	}

	LoginRequest struct {
		Username string `json:"username" validate:"required,min=6,max=32"`
		Password string `json:"password" validate:"required,min=6"`
	}

	LoginResponse struct {
		Username     string `json:"username"`
		FullName     string `json:"full_name"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (l RegisterRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

func (l LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
