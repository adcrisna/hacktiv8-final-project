package routes

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
)

func RouterServer() *gin.Engine {
	router := gin.Default()
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login")
	router.PUT("/users/:userId")
	router.DELETE("/users/:userId")

	return router
}
