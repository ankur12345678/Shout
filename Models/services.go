package models

import (
	"gorm.io/gorm"
)

func InitUserRepo(DB *gorm.DB) IUser {
	return &userRepo{
		DB: DB,
	}
}

func InitPostRepo(DB *gorm.DB) IPost {
	return &postRepo{
		DB: DB,
	}
}

func InitCommentRepo(DB *gorm.DB) IComment {
	return &commentRepo{
		DB: DB,
	}
}
