package models

import (
	"time"
)

type Comment struct {
	ID        int       `gorm:"primaryKey;column:comment_id" json:"id"`
	UserID    int       `gorm:"not null;type:int" json:"user_id" validate:"required"`
	PhotoID   int       `gorm:"not null;type:int" json:"photo_id" validate:"required"`
	Message   string    `gorm:"not null;type:varchar(255)" json:"message" validate:"required"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at"`
	Users     User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Photo     Photo     `gorm:"foreignKey:PhotoID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
