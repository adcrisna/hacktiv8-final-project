package models

import (
	"errors"
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey;column:user_id" json:"id"`
	Username  string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"username" valid:"required~ Your username is required"`
	Email     string    `gorm:"not null;uniqueIndex;type:varchar(255)" json:"email" form:"email" valid:"required~Your full email is required,email~Invalid email format"`
	Password  string    `gorm:"not null;type:varchar(255)" json:"password" form:"password" valid:"required~Your full password is required, minstringlength(6)~Password as to have a minimum length of 6 chaarcters"`
	Age       int       `gorm:"not null;type:int" json:"age" valid:"required~Your Age is required"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(u)
	u.Password = helpers.HashPassword(u.Password)
	if u.Age <= 8 {
		return errors.New("your age must be greater than 8")
	}
	return errCreate
}

type CreateUser struct {
	ID        int        `json:"id"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	Username  string     `json:"username"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
