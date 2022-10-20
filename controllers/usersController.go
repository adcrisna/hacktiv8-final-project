package controllers

import (
	"final-project/models"
	"final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: *service,
	}
}

func (u *UserController) RegisterUser(c *gin.Context) {
	var req models.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.UserRegister(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) LoginUser(c *gin.Context) {
	var req models.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.LoginUser(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var req models.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	req.ID = userId
	result := u.userService.UpdateUser(id, req)

	c.JSON(result.Status, result.Payload)

}

func (u *UserController) DeleteUser(c *gin.Context) {
	var req models.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
	}

	req.ID = userId
	result := u.userService.DeleteUser(id, req)
	c.JSON(result.Status, result.Payload)
}
