package app

import (
	postgresql_db "maranatha_web/datasources/postgresql"
	redis_db "maranatha_web/datasources/redis"
)

func StartApplication() {

	postgresql_db.GetDB()
	redis_db.GetRedisClient()
	r := SetupRouter()

	err := r.Run("127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

}
