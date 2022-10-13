package models

type SocialMedia struct {
	ID             int    `gorm:"primaryKey;column:social_media_id" json:"id" validate:"required"`
	Name           string `gorm:"not null;type:varchar(255)" json:"name" validate:"required"`
	SocialMediaUrl string `gorm:"not null;type:varchar(255)" json:"social_media_url" validate:"required"`
	UserID         int    `gorm:"not null;type:int" json:"user_id" validate:"required"`
	Users          User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
