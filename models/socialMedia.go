package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int    `gorm:"primaryKey;column:social_media_id" json:"id"`
	Name           string `gorm:"not null;type:varchar(255)" json:"name" valid:"required~ Your name is required"`
	SocialMediaUrl string `gorm:"not null;type:varchar(255)" json:"social_media_url" valid:"required~ Your social media url is required"`
	UserID         int    `gorm:"not null;type:int" json:"user_id"`
	Users          User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
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
