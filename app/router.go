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
			protected.DELETE("/delete-book/", controllers.DeleteBook)

			//Prayer_request Routes
			protected.POST("/create-prayer/", controllers.CreatPrayerRequest)
			protected.GET("/get-author-prayer/", controllers.GetAllPrayerRequestsByAuthor)
			protected.GET("/get-all-prayers/", controllers.GetAllPrayerRequests)
			protected.DELETE("/delete-prayer/", controllers.DeletePrayerRequest)
			protected.GET("/get-single-prayer/", controllers.GetSinglePrayerRequest)
			protected.PUT("/update-prayer/", controllers.UpdatePrayerRequest)

			//Testimonies Routes
			protected.POST("/create-testimonies/", controllers.CreateTestimony)
			protected.GET("/get-author-testimonies/", controllers.GetAllTestimoniesByAuthor)
			protected.GET("/get-testimonies/", controllers.GetAllTestimonies)
			protected.DELETE("/delete-testimony/", controllers.DeleteTestimony)
			protected.GET("/get-single-testimony/", controllers.GetSingleTestimony)
			protected.PUT("/update-testimony/", controllers.UpdateTestimony)

			//Events Routes
			protected.POST("/create-event/", controllers.CreatEventsPost)
			protected.GET("/get-author-events/", controllers.GetAllEventsByAuthor)
			protected.GET("/get-events/", controllers.GetAllEvents)
			protected.DELETE("/delete-event/", controllers.DeleteEvent)
			protected.GET("/get-single-event/", controllers.GetSingleEvent)
			protected.PUT("/update-event/", controllers.UpdateEventsPost)

			//Genre Routes
			protected.POST("/create-genre/", controllers.CreatGenrePost)
			protected.GET("/get-genres/", controllers.GetAllGenres)
			protected.DELETE("/delete-genre/", controllers.DeleteGenre)
			protected.GET("/get-single-genre/", controllers.GetSingleGenre)
			protected.PUT("/update-genre/", controllers.UpdateGenre)

			//Jobs Routes
			protected.POST("/create-event-job/", controllers.CreateEventJob)
			protected.GET("/get-event-jobs/", controllers.GetAllEventJobs)
			protected.GET("/get-job-by-event/", controllers.GetJobByEvent)
			protected.DELETE("/delete-job/", controllers.DeleteJob)
			protected.PUT("/update-event-job/", controllers.UpdateJob)

			//Volunteer Job Routes
			protected.POST("/sub-event-job/", controllers.SubscribeToEventJob)
			//protected.GET("/get-event-jobs/", controllers.GetAllEventJobs)
			//protected.GET("/get-job-by-event/", controllers.GetJobByEvent)
			//protected.DELETE("/delete-job/", controllers.DeleteJob)
			//protected.PUT("/update-event-job/", controllers.UpdateJob)
		}
	}

	return router
}
