package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `gorm:"primaryKey;column:photo_id" json:"id"`
	Title     string    `gorm:"not null;type:varchar(255)" json:"title" valid:"required~ Your title is required"`
	Caption   string    `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl  string    `gorm:"not null;type:varchar(255)" json:"photo_url" valid:"required~ Your Photo URL is required"`
	UserID    int       `gorm:"not null;type:int" json:"user_id"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at"`
	Users     User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	return errCreate
}

type CreatePhoto struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type GetAllPhotos struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdateAt  time.Time  `json:"updated_at,omitempty"`
	User      *UserPhoto `json:"user,omitempty"`
}

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
