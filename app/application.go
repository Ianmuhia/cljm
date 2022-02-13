package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"maranatha_web/internal/config"
	"maranatha_web/internal/controllers"
	"maranatha_web/internal/controllers/token"
	"maranatha_web/internal/datasources/fcm_client"
	"maranatha_web/internal/datasources/filestorage"
	mailClient "maranatha_web/internal/datasources/mail"
	redisDb "maranatha_web/internal/datasources/redis"
	"maranatha_web/internal/logger"
	"maranatha_web/internal/repository"
	"maranatha_web/internal/services"
)

var app config.AppConfig

//const jwtKey = "*#*Johnte2536290248"

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error  .env file")
	}
	jwtKey := os.Getenv("JwtSecret")

	//configure logger
	zl := logger.GetLogger()
	app.ErrorLog = zl
	app.InfoLog = zl
	_ = config.NewAppConfig(zl, zl)

	//Create the new token maker
	_, err = token.NewJWTMaker(jwtKey)
	if err != nil {

		return
	}

	//Get database connection
	db := repository.GetDatabaseConnection()
	dao := repository.NewPostgresRepo(db, &app)

	//
	//Get fcm connection
	messagingService := fcm_client.GetFcmConnection()

	//
	//
	//Get file storage connection
	connection, bucketName, minioErr := filestorage.GetMinioConnection()
	if minioErr != nil {
		log.Panicln(err)
	}
	filestorage.NewMinoRepo(connection, &app)
	log.Printf("mino endpoint is %s", connection.EndpointURL())

	fs := filestorage.MinioRepo{
		App:          &app,
		MinioStorage: connection,
	}
	app.StorageURL = connection.EndpointURL()
	app.StorageBucket = bucketName

	//connect to mail server
	mailServer := mailClient.GetMailServer()

	//initiate services
	booksService := services.NewBookService(dao)
	eventsService := services.NewEventsService(dao)
	genresService := services.NewGenreService(dao)
	jobsService := services.NewJobsService(dao)
	newsService := services.NewNewsService(dao)
	partnersService := services.NewChurchPartnersService(dao, &app)
	prayerRequestService := services.NewPrayerRequestService(dao)
	sermonServices := services.NewSermonService(dao)
	testimonyService := services.NewTestimoniesService(dao)
	usersService := services.NewUsersService(dao)
	volunteerJobService := services.NewVolunteerChurchJobService(dao)
	dailyVerseService := services.NewDailyVerseService()
	mailService := services.NewMailService(mailServer)
	fcmService := services.NewFcmService(messagingService)

	redisDb.GetRedisClient()

	allServices := controllers.NewRepo(
		mailService,
		fcmService,
		&app,
		booksService,
		&fs,
		dailyVerseService,
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

	r := SetupRouter()

	err = r.Run("localhost:8090")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
