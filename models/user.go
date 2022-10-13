package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"username" validate:"required"`
	Email     string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"email" vvalidate:"required"`
	Password  string    `gorm:"not null;type:varchar(255)" json:"password" validate:"required"`
	Age       int       `gorm:"not null;type:int" json:"age" validate:"required"`
	CreateAt  time.Time `gorm:"not null;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at"`
}
