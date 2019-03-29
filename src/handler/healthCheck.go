package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// HealthCheck manual update db
func HealthCheck(c *gin.Context) {
	log.Println("pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
