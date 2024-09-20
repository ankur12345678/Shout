package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `gorm:"primaryKey" json:"-"`
	UserUUID       string         `gorm:"unique;not null;" json:"user_uuid"`
	Name           string         `gorm:"not null" json:"name"`
	Email          string         `gorm:"unique" json:"email"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	UserName       string         `gorm:"unique;not null;" json:"username"`
	ProfilePicture string         `json:"profile_picture"`
	Password       string         `gorm:"not null" json:"omit"`
}

type userRepo struct {
	DB *gorm.DB
}

// Create implements IUser.
func (r *userRepo) Create(u *User) error {
	return r.CreateWithTx(r.DB, u)
}

// CreateWithTx implements IUser.
func (r *userRepo) CreateWithTx(tx *gorm.DB, u *User) error {
	err := tx.Model(&User{}).Create(&u).Error
	return err
}

// Delete implements IUser.
func (r *userRepo) Delete(UserUUID string) error {
	err := r.DB.Where(&User{UserUUID: UserUUID}).Delete(&User{}).Error
	return err
}

// GetWithTx implements IUser.
func (r *userRepo) GetWithTx(where *User, tx *gorm.DB) (*User, error) {
	var user User
	err := tx.Model(&User{}).Where(where).First(&user).Error
	return &user, err
}

// Update implements IUser.
func (r *userRepo) Update(u *User, UserUUID string) error {
	return r.UpdateWithTx(r.DB, u, UserUUID)
}

// UpdateWithTx implements IUser.
func (r *userRepo) UpdateWithTx(tx *gorm.DB, u *User, UserUUID string) error {
	err := tx.Model(&User{}).Where(&User{UserUUID: UserUUID}).Updates(u).Error
	return err
}

func (r *userRepo) GetById(UserUUID string) (*User, error) {
	return r.GetWithTx(&User{UserUUID: UserUUID}, r.DB)
}

func (r *userRepo) GetByEmail(Email string) (*User, error) {
	return r.GetWithTx(&User{Email: Email}, r.DB)
}

func (r *userRepo) GetByUserName(UserName string) (*User, error) {
	return r.GetWithTx(&User{UserName: UserName}, r.DB)
}
