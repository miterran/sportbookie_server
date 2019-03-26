package handler

import (
	"sport_bookie_server/src/result"
	"github.com/gin-gonic/gin"
	"sport_bookie_server/src/scheduler"
	"log"
)


// CheckHandler manual update db
func CheckHandler(c *gin.Context) {
	scheduler.UpdateGame()
	scheduler.UpdateScore()
	result.CheckOpenBet()
	log.Println("done")
}