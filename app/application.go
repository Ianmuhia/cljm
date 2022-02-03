package app

import (
	"log"
	"maranatha_web/internal/controllers"
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
	services.NewBookService(dao)
	services.NewEventsService(dao)
	services.NewGenreService(dao)
	services.NewJobsService(dao)
	services.NewNewsService(dao)
	services.NewChurchPartnersService(dao)
	services.NewPrayerRequestService(dao)
	services.NewSermonService(dao)
	services.NewTestimoniesService(dao)
	services.NewUsersService(dao)
	services.NewVolunteerChurchJobService(dao)
	//

	cc := controllers.NewRepo(&app, dao)
	controllers.NewHandlers(cc)

	//Get file storage connection
	connection, minioErr := filestorage.GetMinioConnection()
	if minioErr != nil {
		log.Panicln(err)
	}
	filestorage.NewMinoRepo(connection, &app)
	log.Printf("mino endpoint is %s", connection.EndpointURL())
	//redisDb.GetRedisClient()
	//mailClient.GetMailServer()

	r := SetupRouter()

	err = r.Run("127.0.0.1:8090")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
