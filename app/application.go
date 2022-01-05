package app

import (
	postgresql_db "maranatha_web/datasources/postgresql"
)

func StartApplication() {

	postgresql_db.GetDB()

	r:= SetupRouter()
	
	err := r.Run("127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

}
