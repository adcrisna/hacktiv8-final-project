package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialmedia *models.SocialMedia) (*models.SocialMedia, error)
	GetAllSocialMedias() (*[]models.SocialMedia, error)
	GetSocialMediaByID(socialMediaId int) (*models.SocialMedia, error)
	UpdateSocialMedia(socialMediaId int, socialmedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMedia(socialMediaId int) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db}
}

func (s *socialMediaRepository) CreateSocialMedia(socialmedia *models.SocialMedia) (*models.SocialMedia, error) {
	return socialmedia, s.db.Create(&socialmedia).Error
}

func (s *socialMediaRepository) GetAllSocialMedias() (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := s.db.Preload("users").Find(&socialMedias).Error
	return &socialMedias, err
}

func (s *socialMediaRepository) GetSocialMediaByID(socialMediaId int) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	err := s.db.Preload("User").Where("user_id=?", socialMediaId).Find(&socialMedia).Error
	return &socialMedia, err
}

func (s *socialMediaRepository) UpdateSocialMedia(socialMediaId int, socialMediaUpdate *models.SocialMedia) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	err := s.db.Model(&socialMedia).Clauses(clause.Returning{}).Where("user_id=?", socialMediaId).Updates(socialMediaUpdate).Error
	return &socialMedia, err
}

func (s *socialMediaRepository) DeleteSocialMedia(socialMediaId int) error {
	var socialMedia models.SocialMedia

	err := s.db.Where("user_id=?", socialMediaId).Delete(&socialMedia).Error
	return err
}
