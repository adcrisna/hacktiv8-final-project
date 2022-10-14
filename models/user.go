package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;column:user_id" json:"-"`
	Username  string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"username" validate:"required" binding:"required"`
	Email     string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"email" validate:"required" binding:"required"`
	Password  string    `gorm:"not null;type:varchar(255)" json:"password" validate:"required" binding:"required"`
	Age       int       `gorm:"not null;type:int" json:"age" validate:"required" binding:"required"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at"`
}

type UserResponse struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
