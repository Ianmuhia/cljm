package app

import (
	"fmt"
	"log"

	"go.uber.org/zap" //nolint:goimports
	"maranatha_web/controllers/token"
	mail_client "maranatha_web/datasources/mail"
	postgresql_db "maranatha_web/datasources/postgresql"
	redis_db "maranatha_web/datasources/redis"
)

type Config struct {
	DC *zap.Logger
}

const jwtKey = "*#*Johnte2536290248"

func StartApplication() {
	maker, err := token.NewJWTMaker(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(maker)
	//postgresql_db.GetDB()
	postgresql_db.GetBunDB()

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

}
