package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int `gorm:"primaryKey;column:comment_id" json:"id"`
	UserID    int `json:"user_id"`
	User      *User
	PhotoID   int `json:"photo_id"`
	Photo     *Photo
	Message   string    `gorm:"not null;" json:"message" valid:"required~ Message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	return errCreate
}

type CreateComment struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type GetAllCommentsWithUserAndPhoto struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	UpdateAt  time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      *UserComment
	Photo     *PhotoComment
}

type UserComment struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}
