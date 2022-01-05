package app

import (
	"maranatha_web/controllers/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	router.SetTrustedProxies([]string{"0.0.0.0", "localhost"})

	// Middlewares
	// router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/users/login/", users_controller.Login)
			public.POST("/users/register/", users_controller.RegisterUser)

		}

		// here
		// protected := api.Group("/protected").Use(middlewares.Authz())
		// {
		// 	protected.GET("/profile", controllers.Profile)
		// }
	}

	return router
}
