package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint           `gorm:"primaryKey" json:"-"`
	CommentUUID string         `gorm:"unique;not null;" json:"comment_uuid"`
	UserID      int            `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	PostID      int            `gorm:"not null" json:"post_id"`
	Post        Post           `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE;"`
	Content     string         `json:"content"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type commentRepo struct {
	DB *gorm.DB
}

// Create implements IUser.
func (r *commentRepo) Create(c *Comment) error {
	return r.CreateWithTx(r.DB, c)
}

// CreateWithTx implements IUser.
func (r *commentRepo) CreateWithTx(tx *gorm.DB, c *Comment) error {
	err := tx.Model(&Comment{}).Create(&c).Error
	return err
}

// Delete implements IUser.
func (r *commentRepo) Delete(CommentUUID string) error {
	err := r.DB.Where(&Comment{CommentUUID: CommentUUID}).Delete(&Comment{}).Error
	return err
}

// GetWithTx implements IUser.
func (r *commentRepo) GetWithTx(where *Comment, tx *gorm.DB) (*Comment, error) {
	var comment Comment
	err := tx.Model(&Comment{}).Where(where).First(&comment).Error
	return &comment, err
}

// Update implements IUser.
func (r *commentRepo) Update(c *Comment, CommentUUID string) error {
	return r.UpdateWithTx(r.DB, c, CommentUUID)
}

// UpdateWithTx implements IUser.
func (r *commentRepo) UpdateWithTx(tx *gorm.DB, c *Comment, CommentUUID string) error {
	err := tx.Model(&Comment{}).Where(&Comment{CommentUUID: CommentUUID}).Updates(c).Error
	return err
}

func (r *commentRepo) GetById(CommentUUID string) (*Comment, error) {
	return r.GetWithTx(&Comment{CommentUUID: CommentUUID}, r.DB)
}
