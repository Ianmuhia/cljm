package app

import (
	users_controller "maranatha_web/controllers"
	"maranatha_web/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"0.0.0.0", "localhost"}) //nolint:errcheck

	// Middlewares
	// router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/users/login/", users_controller.Login)
			public.POST("/users/register/", users_controller.RegisterUser)
			public.POST("/users/verifyemail/", users_controller.VerifyEmailCode)

		}

		// here
		protected := api.Group("/protected").Use(middleware.CORSMiddleware(), middleware.AuthMiddleware())
		{
			protected.GET("/profile/", users_controller.TryAuthMiddlewareMiddleware)
			protected.GET("/update-profile-image/", users_controller.UpdateUserProfileImage)
			protected.POST("/create-news/", users_controller.CreatNewsPost)
			protected.GET("/get-users/", users_controller.GetAllUsers)
			protected.GET("/get-news/", users_controller.GetAllNewsPost)
			protected.DELETE("/delete-news/", users_controller.DeleteNewsPost)
			protected.GET("/get-single-news/", users_controller.GeSingleNewsPost)
		}
	}

	return router
}
