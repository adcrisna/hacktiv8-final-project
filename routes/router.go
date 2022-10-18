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

	photoRepo := repositories.NewPhotoRepo(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Auth())
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetAllPhotos)
		photoRouter.PUT("/:photoId", photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", photoController.DeletePhoto)
	}

	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo)
	commentContoller := controllers.NewCommentContoller(commentService)

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Auth())
		commentRouter.POST("/", commentContoller.CreateComment)
		commentRouter.GET("/", commentContoller.GetAllComments)
		commentRouter.PUT("/:commentId", commentContoller.UpdateComment)
		commentRouter.DELETE("/:commentId", commentContoller.DeleteComment)

	}

	socialMediaRepo := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Auth())
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", socialMediaController.GetAllSocialMedias)
		socialMediaRouter.PUT("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", socialMediaController.DeleteSocialMedia)
	}

	return router
}
