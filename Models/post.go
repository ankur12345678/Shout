package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	PostUUID     string         `gorm:"unique;not null;" json:"post_uuid"`
	Title        string         `gorm:"not null" json:"title"`
	Content      string         `gorm:"not null" json:"content"`
	UserID       int            `gorm:"not null" json:"user_id"`
	User         User           `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	Likes        int            `json:"Likes"`
	Replies      int            `json:"replies"`
	ParentPostId *int           `gorm:"index;constraints:OnDelete:CASCADE" json:"parent_post_id"`
	ParentPost   *Post          `gorm:"foreignKey:ParentPostId;references:ID;constraints:OnDelete:CASCADE;" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-" `
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
func (r *postRepo) Delete(PostID uint) error {
	err := r.DB.Where(&Post{ID: PostID}).Delete(&Post{}).Error
	return err
}

// GetWithTx implements IUser.
func (r *postRepo) GetWithTx(where *Post, tx *gorm.DB) (*Post, error) {
	var post Post
	err := tx.Model(&Post{}).Where(where).First(&post).Error
	return &post, err
}

// Update implements IUser.
func (r *postRepo) Update(p *Post, PostID uint) error {
	return r.UpdateWithTx(r.DB, p, PostID)
}

// UpdateWithTx implements IUser.
func (r *postRepo) UpdateWithTx(tx *gorm.DB, p *Post, PostID uint) error {
	err := tx.Model(&Post{}).Where(&Post{ID: PostID}).Updates(p).Error
	return err
}

func (r *postRepo) GetById(PostID uint) (*Post, error) {
	return r.GetWithTx(&Post{ID: PostID}, r.DB)
}
