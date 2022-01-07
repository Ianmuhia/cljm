package app

import (
	// "go/token"
	"maranatha_web/controllers/token"
	postgresql_db "maranatha_web/datasources/postgresql"
	redis_db "maranatha_web/datasources/redis"

	"go.uber.org/zap"
)

type Config struct {
	DC *zap.Logger
}
const jwtKey = "Johnte"

func StartApplication() {
	token.NewJWTMaker(jwtKey)
	postgresql_db.GetDB()
	redis_db.GetRedisClient()
	//lovgg := logger.GetLogger()
	//Config{
	//	DC: lovgg,
	//}
	r := SetupRouter()

	err := r.Run("127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

}
