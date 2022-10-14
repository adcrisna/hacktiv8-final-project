package controllers

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func RegisterUser(ctx *gin.Context) {
	db := database.Connection()
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !valid(user.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bukan email"})
		return
	}
	if user.Age <= 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "minimal umur untuk mendaftar 8 tahun"})
		return
	}
	if len(user.Password) <= 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "minimal panjang 6 karakter"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newUser := models.User{Username: user.Username, Email: user.Email, Password: string(hashPassword), Age: user.Age}
	userRes := models.UserResponse{}
	if err := db.Create(&newUser).Scan(&userRes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "create data user success", "data": userRes})

}
