package app

import (
	"log"

	"maranatha_web/controllers/token"
	mail_client "maranatha_web/datasources/mail"
	"maranatha_web/datasources/minio"
	redis_db "maranatha_web/datasources/redis"
)

const jwtKey = "*#*Johnte2536290248"

func StartApplication() {
	_, err := token.NewJWTMaker(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}
	//postgresql_db.GetBunDB()
	connection, err := minio.MinioConnection()

	if err != nil {
		log.Panicln(err)
	}
	log.Printf("mino endpoint is %s", connection.EndpointURL())
	redis_db.GetRedisClient()
	//task_client.GetTasksClient()
	mail_client.GetMailServer()

	r := SetupRouter()

	err = r.Run("127.0.0.1:8090")
	if err != nil {
		panic(err)
	}
	//worker.GetTasksWorker()
	//task_client.GetTasksClient()
	//
}
