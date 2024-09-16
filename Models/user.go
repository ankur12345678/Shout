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
