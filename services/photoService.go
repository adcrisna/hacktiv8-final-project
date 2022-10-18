package services

import (
	"final-project/models"
	"final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepository
}

func NewPhotoService(repo repositories.PhotoRepository) *PhotoService {
	return &PhotoService{
		photoRepo: repo,
	}
}

var allPhotos []models.GetAllPhotos

func (p *PhotoService) CreatePhoto(request models.CreatePhoto) *models.Response {

	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserID:   request.UserID,
	}

	photoData, err := p.photoRepo.CreatePhoto(&photo)

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &models.Response{
		Status: http.StatusCreated,
		Payload: &models.GetAllPhotos{
			ID:        photoData.ID,
			Title:     photoData.Title,
			Caption:   photoData.Caption,
			PhotoUrl:  photoData.PhotoUrl,
			UserID:    photoData.UserID,
			CreatedAt: photoData.CreatedAt,
		},
	}
}

func (p *PhotoService) GetAllPhotos() *models.Response {
	result, err := p.photoRepo.GetAllPhotos()

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	allPhotos = nil
	for _, p := range *result {
		allPhotos = append(allPhotos, models.GetAllPhotos{
			ID:        p.ID,
			Title:     p.Title,
			Caption:   p.Caption,
			PhotoUrl:  p.PhotoUrl,
			UserID:    p.UserID,
			CreatedAt: p.CreatedAt,
			UpdateAt:  p.UpdatedAt,
			User: &models.UserPhoto{
				Email:    p.User.Email,
				Username: p.User.Username,
			},
		})
	}

	return &models.Response{
		Status:  http.StatusOK,
		Payload: allPhotos,
	}
}

func (p *PhotoService) UpdatePhoto(request *models.CreatePhoto) *models.Response {
	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
	}

	checkData, err := p.photoRepo.GetPhotoByID(request.ID)

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkData.ID < 1 {
		return &models.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "record not found",
			},
		}
	}

	if request.UserID != checkData.UserID {
		return &models.Response{
			Status: http.StatusForbidden,
			Payload: gin.H{
				"error": "forbidden - only owner can update data",
			},
		}
	}

	updatePhoto, err := p.photoRepo.UpdatePhoto(request.ID, &photo)

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &models.Response{
		Status: http.StatusOK,
		Payload: &models.GetAllPhotos{
			ID:       updatePhoto.ID,
			Title:    updatePhoto.Title,
			Caption:  updatePhoto.Caption,
			PhotoUrl: updatePhoto.PhotoUrl,
			UserID:   updatePhoto.UserID,
			UpdateAt: updatePhoto.UpdatedAt,
		},
	}

}

func (p *PhotoService) DeletePhoto(request *models.CreatePhoto) *models.Response {
	checkData, err := p.photoRepo.GetPhotoByID(request.ID)

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkData.ID < 1 {
		return &models.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "record not found",
			},
		}
	}

	if request.UserID != checkData.UserID {
		return &models.Response{
			Status: http.StatusForbidden,
			Payload: gin.H{
				"error": "forbidden - only owner can delete data",
			},
		}
	}

	err = p.photoRepo.DeletePhoto(request.ID)

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &models.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Your photo has been successfully deleted",
		},
	}

}
