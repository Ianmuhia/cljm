package app

import (
	"github.com/gin-gonic/gin"

	"maranatha_web/internal/controllers"
	"maranatha_web/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	_ = router.SetTrustedProxies([]string{"0.0.0.0", "localhost"})

	api := router.Group("/api")
	{
		public := api.Group("/auth")
		{
			public.POST("/users/login/", controllers.Repo.Login)
			public.POST("/users/register/", controllers.Repo.RegisterUser)
			public.POST("/users/verifyemail/", controllers.Repo.VerifyEmailCode)
			public.POST("/users/forgotpassword/", controllers.Repo.ForgotPassword)
			public.POST("/users/verifypassrstcode/", controllers.Repo.VerifyPassWordResetCode)
			public.POST("/users/newpassword/", controllers.Repo.ResetPassword)

		}

		// here
		protected := api.Group("/protected").Use(middleware.AuthMiddleware())
		// protected := api.Group("/protected")
		{
			///Users Routes
			//protected.GET("/profile/", controllers.Repo.TryAuthMiddlewareMiddleware)
			protected.POST("/update-profile-image/", controllers.Repo.UpdateUserProfileImage)
			protected.GET("/get-users/", controllers.Repo.GetAllUsers)
			protected.POST("/update-password/", controllers.Repo.UpdateUserPassword)

			///News Routes
			protected.POST("/create-news/", controllers.Repo.CreatNewsPost)
			protected.POST("/get-author-news/", controllers.Repo.GetAllNewsPostByAuthor)
			protected.GET("/get-news/", controllers.Repo.GetAllNewsPost)
			protected.DELETE("/delete-news/", controllers.Repo.DeleteNewsPost)
			protected.GET("/get-single-news/", controllers.Repo.GeSingleNewsPost)
			protected.PUT("/update-news/", controllers.Repo.UpdateNewsPost)

			///Church Partners Routes
			protected.POST("/create-partner/", controllers.Repo.CreateChurchPartner)
			protected.GET("/get-partner/", controllers.Repo.GetSingleChurchPartner)
			protected.DELETE("/delete-partner/", controllers.Repo.DeleteChurchPartner)
			protected.GET("/get-partners/", controllers.Repo.GetAllChurchPartner)
			protected.PUT("/update-partner/", controllers.Repo.UpdateChurchPartner)

			///Sermon Routes
			protected.POST("/create-sermon/", controllers.Repo.CreateSermon)
			protected.GET("/get-sermon/", controllers.Repo.GetSingleSermon)
			protected.DELETE("/delete-sermon/", controllers.Repo.DeleteSermon)
			protected.GET("/get-sermons/", controllers.Repo.GetAllSermons)
			protected.PUT("/update-sermon/", controllers.Repo.UpdateSermon)

			///Books Routes
			protected.POST("/create-book/", controllers.Repo.CreateBook)
			protected.GET("/get-book/", controllers.Repo.GetSingleBookPost)
			protected.PUT("/update-book/", controllers.Repo.UpdateBook)
			protected.GET("/get-books/", controllers.Repo.GetAllBooksPost)
			protected.DELETE("/delete-book/", controllers.Repo.DeleteBook)

			//Prayer_request Routes
			protected.POST("/create-prayer/", controllers.Repo.CreatPrayerRequest)
			protected.GET("/get-author-prayer/", controllers.Repo.GetAllPrayerRequestsByAuthor)
			protected.GET("/get-all-prayers/", controllers.Repo.GetAllPrayerRequests)
			protected.DELETE("/delete-prayer/", controllers.Repo.DeletePrayerRequest)
			protected.GET("/get-single-prayer/", controllers.Repo.GetSinglePrayerRequest)
			protected.PUT("/update-prayer/", controllers.Repo.UpdatePrayerRequest)

			//Testimonies Routes
			protected.POST("/create-testimonies/", controllers.Repo.CreateTestimony)
			protected.GET("/get-author-testimonies/", controllers.Repo.GetAllTestimoniesByAuthor)
			protected.GET("/get-testimonies/", controllers.Repo.GetAllTestimonies)
			protected.DELETE("/delete-testimony/", controllers.Repo.DeleteTestimony)
			protected.GET("/get-single-testimony/", controllers.Repo.GetSingleTestimony)
			protected.PUT("/update-testimony/", controllers.Repo.UpdateTestimony)

			//Events Routes
			protected.POST("/create-event/", controllers.Repo.CreatEventsPost)
			protected.GET("/get-author-events/", controllers.Repo.GetAllEventsByAuthor)
			protected.GET("/get-events/", controllers.Repo.GetAllEvents)
			protected.DELETE("/delete-event/", controllers.Repo.DeleteEvent)
			protected.GET("/get-single-event/", controllers.Repo.GetSingleEvent)
			protected.PUT("/update-event/", controllers.Repo.UpdateEventsPost)

			//Genre Routes
			protected.POST("/create-genre/", controllers.Repo.CreatGenrePost)
			protected.GET("/get-genres/", controllers.Repo.GetAllGenres)
			protected.DELETE("/delete-genre/", controllers.Repo.DeleteGenre)
			protected.GET("/get-single-genre/", controllers.Repo.GetSingleGenre)
			protected.PUT("/update-genre/", controllers.Repo.UpdateGenre)

			//Jobs Routes
			protected.POST("/create-event-job/", controllers.Repo.CreateEventJob)
			protected.GET("/get-event-jobs/", controllers.Repo.GetAllEventJobs)
			protected.GET("/get-job-by-event/", controllers.Repo.GetJobByEvent)
			protected.DELETE("/delete-job/", controllers.Repo.DeleteJob)
			protected.PUT("/update-event-job/", controllers.Repo.UpdateJob)

			//Daily Scriptures Routes
			protected.POST("/add-daily-scriptures/", controllers.Repo.GetDailyScriptures)

			//Volunteer Job Routes
			protected.POST("/sub-event-job/", controllers.Repo.SubscribeToEventJob)
			protected.GET("/get-user-volunteered-jobs/", controllers.Repo.GetUserVolunteeredJobs)

		}
	}

	return router
}
