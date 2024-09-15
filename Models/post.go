package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	PostUUID  string `gorm:"unique"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
