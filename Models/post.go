package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	PostUUID   string         `gorm:"unique;not null;" json:"post_uuid"`
	Title      string         `gorm:"not null" json:"title"`
	Content    string         `gorm:"not null" json:"content"`
	UserID     int            `gorm:"not null" json:"user_id"`
	User       User           `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Reputation int            `json:"reputation"`
	Share      int            `json:"share"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type postRepo struct {
	DB *gorm.DB
}

// Create implements IUser.
func (r *postRepo) Create(p *Post) error {
	return r.CreateWithTx(r.DB, p)
}

// CreateWithTx implements IUser.
func (r *postRepo) CreateWithTx(tx *gorm.DB, p *Post) error {
	err := tx.Model(&Post{}).Create(&p).Error
	return err
}

// Delete implements IUser.
func (r *postRepo) Delete(PostUUID string) error {
	err := r.DB.Where(&Post{PostUUID: PostUUID}).Delete(&Post{}).Error
	return err
}

// GetWithTx implements IUser.
func (r *postRepo) GetWithTx(where *Post, tx *gorm.DB) (*Post, error) {
	var post Post
	err := tx.Model(&Post{}).Where(where).First(&post).Error
	return &post, err
}

// Update implements IUser.
func (r *postRepo) Update(p *Post, PostUUID string) error {
	return r.UpdateWithTx(r.DB, p, PostUUID)
}

// UpdateWithTx implements IUser.
func (r *postRepo) UpdateWithTx(tx *gorm.DB, p *Post, PostUUID string) error {
	err := tx.Model(&Post{}).Where(&Post{PostUUID: PostUUID}).Updates(p).Error
	return err
}

func (r *postRepo) GetById(PostUUID string) (*Post, error) {
	return r.GetWithTx(&Post{PostUUID: PostUUID}, r.DB)
}
