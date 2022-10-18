package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int    `gorm:"primaryKey;column:social_media_id" json:"id"`
	Name           string `gorm:"not null;" json:"name" valid:"required~ Your name is required"`
	SocialMediaUrl string `gorm:"not null;" json:"social_media_url" valid:"required~ Your social media url is required"`
	UserID         int
	User           *User
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	return errCreate
}

type CreateSocialMedia struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type GetAllSocialMedias struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	User           *UserSocialMedia
}

type UserSocialMedia struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
}
