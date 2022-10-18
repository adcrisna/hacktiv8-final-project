package services

import (
	"final-project/models"
	"final-project/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepository
}

func NewSocialMediaService(repo repositories.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		socialMediaRepo: repo,
	}
}

var allSocialMedias []models.GetAllSocialMedias

func (s *SocialMediaService) CreateSocialMedia(request models.CreateSocialMedia) *models.Response {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserID:         request.UserID,
	}

	socialMediaData, err := s.socialMediaRepo.CreateSocialMedia(&socialMedia)

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
		Payload: models.CreateSocialMedia{
			ID:             socialMediaData.ID,
			Name:           socialMediaData.Name,
			SocialMediaUrl: socialMediaData.SocialMediaUrl,
			UserID:         socialMediaData.UserID,
			CreatedAt:      socialMediaData.CreatedAt,
		},
	}

}

func (s *SocialMediaService) GetAllSocialMedias() *models.Response {
	allSocialMedias = nil
	socialMediaDatas, err := s.socialMediaRepo.GetAllSocialMedias()

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	for _, v := range *socialMediaDatas {
		allSocialMedias = append(allSocialMedias, models.GetAllSocialMedias{
			ID:             v.ID,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserID:         v.UserID,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User: &models.UserSocialMedia{
				ID:       v.User.ID,
				Username: v.User.Username,
			},
		})
	}

	return &models.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"social_medias": allSocialMedias,
		},
	}

}

func (s *SocialMediaService) UpdateSocialMedia(request models.CreateSocialMedia) *models.Response {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	checkData, err := s.socialMediaRepo.GetSocialMediaByID(request.ID)

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

	updateData, err := s.socialMediaRepo.UpdateSocialMedia(request.ID, &socialMedia)

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
		Payload: models.CreateSocialMedia{
			ID:             updateData.ID,
			Name:           updateData.Name,
			SocialMediaUrl: updateData.SocialMediaUrl,
			UserID:         updateData.UserID,
			CreatedAt:      updateData.CreatedAt,
			UpdatedAt:      updateData.UpdatedAt,
		},
	}
}

func (s *SocialMediaService) DeleteSocialMedia(request models.CreateSocialMedia) *models.Response {
	checkData, err := s.socialMediaRepo.GetSocialMediaByID(request.ID)

	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkData.ID < 1 || checkData == nil {
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

	err = s.socialMediaRepo.DeleteSocialMedia(request.ID)

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
			"message": "Your social media has been successfully deleted",
		},
	}
}
