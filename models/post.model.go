package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title    string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content  string    `gorm:"not null" json:"content,omitempty"`
	Image    string    `gorm:"not null" json:"image,omitempty"`
	User     uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreateAt time.Time `gorm:"not null" json:"create_at,omitempty"`
	UpdateAt time.Time `gorm:"not null" json:"update_at,omitempty"`
}

type CreatePostRequest struct {
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Image    string    `json:"image" binding:"required"`
	User     string    `json:"user,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}

type UpdatePost struct {
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	Image    string    `json:"image,omitempty"`
	User     string    `json:"user,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}
