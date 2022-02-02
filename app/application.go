package app

import (
	"log"
	"maranatha_web/internal/controllers"
	"maranatha_web/internal/logger"
	"maranatha_web/internal/services"

	"maranatha_web/internal/config"
	"maranatha_web/internal/controllers/token"
	"maranatha_web/internal/datasources/minio"
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

	_, err := token.NewJWTMaker(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}
	db := repository.GetDatabaseConnection()
	dao := repository.NewPostgresRepo(db, &app)
	services.NewBookService(dao)
	cc := controllers.NewRepo(&app, dao)
	controllers.NewHandlers(cc)

	_ = config.NewAppConfig(zl, zl)
	// postgresqlDb.GetDatabaseConnection()

	connection, minioErr := minio.GetMinioConnection()

	if minioErr != nil {
		log.Panicln(err)
	}
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
