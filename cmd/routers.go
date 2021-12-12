package main

import (
	"bitbucket.org/wycemiro/cljm.git/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"0.0.0.0", "localhost"})

	// Middlewares
	router.Use(ErrorHandler)
	router.Use(CORSMiddleware())

	// routes
	router.POST("/register", handlers.Create)
	router.POST("/login", handlers.Login)
	router.GET("/session", handlers.Session)
	router.POST("/createReset", handlers.InitiatePasswordReset)
	router.POST("/resetPassword", handlers.ResetPassword)
	return router
}
