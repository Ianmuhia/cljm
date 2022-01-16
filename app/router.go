package app

import (
	"maranatha_web/controllers"
	"maranatha_web/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	_ = router.SetTrustedProxies([]string{"0.0.0.0", "localhost"})

	api := router.Group("/api")
	{
		public := api.Group("/auth")
		{
			public.POST("/users/login/", controllers.Login)
			public.POST("/users/register/", controllers.RegisterUser)
			public.POST("/users/verifyemail/", controllers.VerifyEmailCode)

		}

		// here
		protected := api.Group("/protected").Use(middleware.CORSMiddleware(), middleware.AuthMiddleware())
		{
			///Users Routes
			protected.GET("/profile/", controllers.TryAuthMiddlewareMiddleware)
			protected.POST("/update-profile-image/", controllers.UpdateUserProfileImage)
			protected.GET("/get-users/", controllers.GetAllUsers)

			///News Routes
			protected.POST("/create-news/", controllers.CreatNewsPost)
			protected.POST("/get-author-news/", controllers.GetAllNewsPostByAuthor)
			protected.GET("/get-news/", controllers.GetAllNewsPost)
			protected.DELETE("/delete-news/", controllers.DeleteNewsPost)
			protected.GET("/get-single-news/", controllers.GeSingleNewsPost)
			protected.PUT("/update-news/", controllers.UpdateNewsPost)

			///Church Partners Routes
			protected.POST("/create-partner/", controllers.CreateChurchPartner)
			protected.GET("/get-partner/", controllers.GetSingleChurchPartner)
			protected.DELETE("/delete-partner/", controllers.DeleteChurchPartner)
			protected.GET("/get-partners/", controllers.GetAllChurchPartner)
			protected.PUT("/update-partner/", controllers.UpdateChurchPartner)

			///Sermon Routes
			protected.POST("/create-sermon/", controllers.CreateSermon)
			protected.GET("/get-sermon/", controllers.GetSingleSermon)
			protected.DELETE("/delete-sermon/", controllers.DeleteSermon)
			protected.GET("/get-sermons/", controllers.GetAllSermons)
			protected.PUT("/update-sermon/", controllers.UpdateSermon)

		}
	}

	return router
}
