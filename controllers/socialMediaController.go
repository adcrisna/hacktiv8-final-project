package controllers

import (
	"final-project/models"
	"final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(service *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		socialMediaService: *service,
	}
}

func (s *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var req models.CreateSocialMedia

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}
	req.UserID = userId

	result := s.socialMediaService.CreateSocialMedia(req)

	c.JSON(result.Status, result.Payload)
}

func (s *SocialMediaController) GetAllSocialMedias(c *gin.Context) {
	result := s.socialMediaService.GetAllSocialMedias()

	c.JSON(result.Status, result.Payload)
}

func (s *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	var req models.CreateSocialMedia

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}
	req.UserID = userId
	req.ID = id

	result := s.socialMediaService.UpdateSocialMedia(req)

	c.JSON(result.Status, result.Payload)

}

func (s *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	var req models.CreateSocialMedia

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}
	req.UserID = userId
	req.ID = id

	result := s.socialMediaService.DeleteSocialMedia(req)

	c.JSON(result.Status, result.Payload)

}
