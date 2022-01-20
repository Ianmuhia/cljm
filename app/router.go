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

			///Books Routes
			protected.POST("/create-book/", controllers.CreateBook)
			protected.GET("/get-book/", controllers.GetSingleBookPost)
			protected.PUT("/update-book/", controllers.UpdateBook)
			protected.GET("/get-books/", controllers.GetAllBooksPost)

			//Prayer_request Routes
			protected.POST("/create-prayer/", controllers.CreatPrayerPost)
			protected.GET("/get-author-prayer/", controllers.GetAllPrayerPostByAuthor)
			protected.GET("/get-prayer/", controllers.GetAllPrayerPost)
			protected.DELETE("/delete-prayer/", controllers.DeletePrayerPost)
			protected.GET("/get-single-prayer/", controllers.GetSinglePrayerPost)
			protected.PUT("/update-prayer/", controllers.UpdatePrayerPost)

			//Testimonies Routes
			protected.POST("/create-testimonies/", controllers.CreatTestimoniesPost)
			protected.POST("/get-author-testimonies/", controllers.GetAllTestimoniesPostByAuthor)
			protected.GET("/get-testimonies/", controllers.GetAllTestimoniesPost)
			protected.DELETE("/delete-testimonies/", controllers.DeleteTestimoniesPost)
			protected.GET("/get-single-testimonies/", controllers.GetSingleTestimoniesPost)
			protected.PUT("/update-testimonies/", controllers.UpdateTestimoniesPost)

			//Events Routes
			protected.POST("/create-events/", controllers.CreatEventsPost)
			protected.POST("/get-author-events/", controllers.GetAllEventsPostByAuthor)
			protected.GET("/get-events/", controllers.GetAllEventsPost)
			protected.DELETE("/delete-events/", controllers.DeleteEventsPost)
			protected.GET("/get-single-events/", controllers.GetSingleEventsPost)
			protected.PUT("/update-events/", controllers.UpdateEventsPost)

			//Genre Routes
			protected.POST("/create-genre/", controllers.CreatGenrePost)
			protected.GET("/get-genre/", controllers.GetAllGenrePost)
			protected.DELETE("/delete-genre/", controllers.DeleteGenrePost)
			protected.GET("/get-single-genre/", controllers.GetSingleGenrePost)
			protected.PUT("/update-genre/", controllers.UpdateGenrePost)
		}
	}

	return router
}
