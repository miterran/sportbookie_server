package router

import (
	"sport_bookie_server/src/handler"
	"sport_bookie_server/src/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run router
func Run(port string) {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", handler.HealthCheck)
	router.GET("/check", handler.ManualUpdate)
	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", middleware.AuthMiddleware.LoginHandler)
	router.Use(middleware.AuthMiddleware.MiddlewareFunc())
	router.POST("/graphql", handler.Graphqlhandler())
	router.GET("/graphql", handler.Graphqlhandler())
	router.Run(port)
}
