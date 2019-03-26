package router

import (
	"fmt"
	"sport_bookie_server/src/handler"
	"sport_bookie_server/src/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run router
func Run(port string) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ping", handler.PingHandler)
	router.GET("/check", handler.CheckHandler)
	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", middleware.AuthMiddleware.LoginHandler)
	router.Use(middleware.AuthMiddleware.MiddlewareFunc())
	router.POST("/graphql", handler.Graphqlhandler())
	router.GET("/graphql", handler.Graphqlhandler())
	router.Run(port)
}
