package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	Username  string    `json:"username" gorm:"unique;" validate:"required,min=6,max=32"`
	FullName  string    `json:"full_name" gorm:"type:text;" validate:"required,min=6,max=32"`
	Password  string    `json:"password" gorm:"type:text;" validate:"required,min=6"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
