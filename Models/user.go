package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `gorm:"primaryKey"`
	UserUUID       string `gorm:"unique"`
	Name           string
	Email          string `gorm:"unique"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserName       string         `gorm:"unique"`
	ProfilePicture string
	Password       string
}
