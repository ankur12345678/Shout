package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint   `gorm:"primaryKey"`
	CommentUUID string `gorm:"unique"`
	UserID      string
	User        User `gorm:"foreignKey:UserID;references:UserUUID;constraint:OnDelete:CASCADE;"`
	PostID      string
	Post        Post `gorm:"foreignKey:PostID;references:PostUUID;constraint:OnDelete:CASCADE;"`
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
