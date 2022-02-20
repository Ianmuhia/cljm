package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"maranatha_web/internal/controllers"
	"maranatha_web/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH, POST, HEAD, GET, OPTIONS , DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// router.Use(middleware.CORSMiddleware())

	_ = router.SetTrustedProxies([]string{"0.0.0.0", "localhost"})

	router.Use(otelgin.Middleware(serviceName))

	api := router.Group("/api")
	{
		api.POST("/batch-upload/", controllers.Repo.BatchUpload)

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

		{
			///Users Routes
			protected.GET("/get-user/", controllers.Repo.GetUser)
			protected.POST("/update-profile-image/", controllers.Repo.UpdateUserProfileImage)
			protected.GET("/get-users/", controllers.Repo.GetAllUsers)
			protected.DELETE("/delete-user/:id", controllers.Repo.DeleteUser)
			protected.POST("/update-password/", controllers.Repo.UpdateUserPassword)
			protected.POST("/update-profile/", controllers.Repo.UpdateUserDetails)

			///News Routes
			protected.POST("/news/", controllers.Repo.CreatNewsPost)
			protected.POST("/get-author-news/", controllers.Repo.GetAllNewsPostByAuthor)
			protected.GET("/news/", controllers.Repo.GetAllNewsPost)
			protected.DELETE("/news/", controllers.Repo.DeleteNewsPost)
			protected.GET("/news-detail/", controllers.Repo.GeSingleNewsPost)
			protected.PUT("/news/", controllers.Repo.UpdateNewsPost)

			///Church Partners Routes
			protected.POST("/partner/", controllers.Repo.CreateChurchPartner)
			protected.GET("/partner/", controllers.Repo.GetSingleChurchPartner)
			protected.DELETE("/partner/", controllers.Repo.DeleteChurchPartner)
			protected.GET("/partners/", controllers.Repo.GetAllChurchPartner)
			protected.PUT("/partner/", controllers.Repo.UpdateChurchPartner)

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

			//Podcast Routes
			protected.POST("/podcast/", controllers.Repo.CreatePodcast)
			protected.POST("/get-author-podcast/", controllers.Repo.GetAllPodcastByAuthor)
			protected.GET("/podcast/", controllers.Repo.GetAllPodcast)
			protected.DELETE("/podcast/", controllers.Repo.DeletePodcast)
			protected.GET("/podcast-detail/", controllers.Repo.GetSinglePodcast)
			protected.PUT("/podcast/", controllers.Repo.UpdatePodcast)

		}
	}

	pprof.Register(router)

	return router
}
