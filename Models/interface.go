package models

import (
	"gorm.io/gorm"
)

type IUser interface {
	GetById(UserUUID string) (*User, error)
	GetByEmail(Email string) (*User, error)
	GetByUserName(UserName string) (*User, error)
	GetWithTx(where *User, tx *gorm.DB) (*User, error)
	Create(u *User) error
	CreateWithTx(tx *gorm.DB, u *User) error
	Update(u *User, UserUUID string) error
	UpdateWithTx(tx *gorm.DB, u *User, UserUUID string) error
	Delete(UserUUID string) error
}

type IPost interface {
	GetById(PostID uint) (*Post, error)
	GetWithTx(where *Post, tx *gorm.DB) (*Post, error)
	Create(p *Post) error
	CreateWithTx(tx *gorm.DB, p *Post) error
	Update(p *Post, PostID uint) error
	UpdateWithTx(tx *gorm.DB, p *Post, PostID uint) error
	Delete(PostID uint) error
}
