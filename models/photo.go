package models

import "time"

type Photo struct {
	ID        int       `gorm:"primaryKey;column:photo_id" json:"id"`
	Title     string    `gorm:"not null;type:varchar(255)" json:"title" validate:"required"`
	Caption   string    `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl  string    `gorm:"not null;type:varchar(255)" json:"photo_url" validate:"required"`
	UserID    int       `gorm:"not null;type:int" json:"user_id" validate:"required"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at"`
	Users     User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
