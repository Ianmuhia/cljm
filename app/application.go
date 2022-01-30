package app

import (
	"encoding/gob"
	"log"

	"maranatha_web/controllers/token"
	mailClient "maranatha_web/datasources/mail"
	"maranatha_web/datasources/minio"
	postgresqlDb "maranatha_web/datasources/postgresql"
	redisDb "maranatha_web/datasources/redis"
	"maranatha_web/models"
)

const jwtKey = "*#*Johnte2536290248"

func StartApplication() {
	gob.Register(models.ChurchPartner{})

	_, err := token.NewJWTMaker(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}

	postgresqlDb.GetDatabaseConnection()

	connection, minioErr := minio.GetMinioConnection()

	if minioErr != nil {
		log.Panicln(err)
	}
	log.Printf("mino endpoint is %s", connection.EndpointURL())
	redisDb.GetRedisClient()
	mailClient.GetMailServer()

	r := SetupRouter()

	err = r.Run("127.0.0.1:8090")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
