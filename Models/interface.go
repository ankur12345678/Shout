package models

import "gorm.io/gorm"

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
	GetById(PostUUID string) (*Post, error)
	GetWithTx(where *Post, tx *gorm.DB) (*Post, error)
	Create(p *Post) error
	CreateWithTx(tx *gorm.DB, p *Post) error
	Update(p *Post, PostUUID string) error
	UpdateWithTx(tx *gorm.DB, p *Post, PostUUID string) error
	Delete(PostUUID string) error
}

type IComment interface {
	GetById(CommentUUID string) (*Comment, error)
	GetWithTx(where *Comment, tx *gorm.DB) (*Comment, error)
	Create(c *Comment) error
	CreateWithTx(tx *gorm.DB, c *Comment) error
	Update(c *Comment, CommentUUID string) error
	UpdateWithTx(tx *gorm.DB, c *Comment, CommentUUID string) error
	Delete(CommentUUID string) error
}
