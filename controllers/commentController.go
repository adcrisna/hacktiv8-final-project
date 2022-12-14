package controllers

import (
	"final-project/models"
	"final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentContoller(service *services.CommentService) *CommentController {
	return &CommentController{
		commentService: service,
	}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var req models.CreateComment

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

	result := cc.commentService.CreateComment(req)

	c.JSON(result.Status, result.Payload)
}

func (cc *CommentController) GetAllComments(c *gin.Context) {
	result := cc.commentService.GetAllComments()

	c.JSON(result.Status, result.Payload)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var req models.CreateComment

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

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	req.UserID = userId
	req.ID = commentId

	result := cc.commentService.UpdateComment(req)
	c.JSON(result.Status, result.Payload)
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	var req models.CreateComment

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
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	req.UserID = userId
	req.ID = commentId

	result := cc.commentService.DeleteComment(req)
	c.JSON(result.Status, result.Payload)
}
