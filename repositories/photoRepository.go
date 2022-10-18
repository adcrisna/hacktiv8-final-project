package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) (*models.Photo, error)
	GetAllPhotos() (*[]models.Photo, error)
	GetPhotoByID(photoId int) (*models.Photo, error)
	UpdatePhoto(photoId int, updatePhoto *models.Photo) (*models.Photo, error)
	DeletePhoto(photoId int) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepository {
	return &photoRepository{db}
}

func (p *photoRepository) CreatePhoto(photo *models.Photo) (*models.Photo, error) {
	return photo, p.db.Create(&photo).Error
}

func (p *photoRepository) GetAllPhotos() (*[]models.Photo, error) {
	var photo []models.Photo
	err := p.db.Preload("User").Find(&photo).Error
	return &photo, err
}

func (p *photoRepository) GetPhotoByID(photoId int) (*models.Photo, error) {
	var photo models.Photo
	err := p.db.Preload("User").Where("user_id=?", photoId).Find(&photo).Error
	return &photo, err
}

func (p *photoRepository) UpdatePhoto(photoId int, updatePhoto *models.Photo) (*models.Photo, error) {
	var photo models.Photo

	err := p.db.Model(&photo).Clauses(clause.Returning{}).Where("user_id=?", photoId).Updates(updatePhoto).Error
	return &photo, err
}

func (p *photoRepository) DeletePhoto(photoId int) error {
	var photo models.Photo

	err := p.db.Where("user_id=?", photoId).Delete(&photo).Error
	return err
}
