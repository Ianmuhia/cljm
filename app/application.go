package app

import (
	"log"
	"maranatha_web/internal/controllers"
	mailClient "maranatha_web/internal/datasources/mail"
	redisDb "maranatha_web/internal/datasources/redis"
	"maranatha_web/internal/logger"
	"maranatha_web/internal/services"

	"maranatha_web/internal/config"
	"maranatha_web/internal/controllers/token"
	"maranatha_web/internal/datasources/filestorage"
	"maranatha_web/internal/repository"
)

var app config.AppConfig

const jwtKey = "*#*Johnte2536290248"

func StartApplication() {
	//configure logger
	zl := logger.GetLogger()
	log.Println(zl)
	app.ErrorLog = zl
	app.InfoLog = zl
	_ = config.NewAppConfig(zl, zl)

	//Create the new token maker
	_, err := token.NewJWTMaker(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}
	//Get database connection
	db := repository.GetDatabaseConnection()
	dao := repository.NewPostgresRepo(db, &app)
	//initiate services
	//
	booksService := services.NewBookService(dao)
	eventsService := services.NewEventsService(dao)
	genresService := services.NewGenreService(dao)
	jobsService := services.NewJobsService(dao)
	newsService := services.NewNewsService(dao)
	partnersService := services.NewChurchPartnersService(dao)
	prayerRequestService := services.NewPrayerRequestService(dao)
	sermonServices := services.NewSermonService(dao)
	testimonyService := services.NewTestimoniesService(dao)
	usersService := services.NewUsersService(dao)
	volunteerJobService := services.NewVolunteerChurchJobService(dao)
	//
	//
	//Get file storage connection
	connection, minioErr := filestorage.GetMinioConnection()
	if minioErr != nil {
		log.Panicln(err)
	}
	filestorage.NewMinoRepo(connection, &app)
	log.Printf("mino endpoint is %s", connection.EndpointURL())

	fs := filestorage.MinioRepo{
		App:          &app,
		MinioStorage: connection,
	}

	allServices := controllers.NewRepo(
		&app,
		booksService,
		&fs,
		eventsService,
		genresService,
		jobsService,
		newsService,
		partnersService,
		prayerRequestService,
		sermonServices,
		testimonyService,
		usersService,
		volunteerJobService)

	controllers.NewHandlers(allServices)

	redisDb.GetRedisClient()
	mailClient.GetMailServer()

	r := SetupRouter()

	err = r.Run("127.0.0.1:8090")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
