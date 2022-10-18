package routes

import (
	"final-project/controllers"
	"final-project/database"
	"final-project/middleware"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func RouterServer() *gin.Engine {
	db := database.Connection()
	router := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewItemService(userRepo)
	userController := controllers.NewUserController(userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.RegisterUser)
		userRouter.POST("/login", userController.LoginUser)
		userRouter.Use(middleware.Auth())
		userRouter.PUT("/:userId", userController.UpdateUser)
		userRouter.DELETE("/:userId", userController.DeleteUser)
	}

	return router
}
