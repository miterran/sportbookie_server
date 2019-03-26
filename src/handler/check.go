package handler

import (
	"log"
	"net/http"
	"sport_bookie_server/src/result"
	"sport_bookie_server/src/scheduler"

	"github.com/gin-gonic/gin"
)

// CheckHandler manual update db
func CheckHandler(c *gin.Context) {
	scheduler.UpdateGame()
	scheduler.UpdateScore()
	result.CheckOpenBet()
	log.Println("done")
	c.JSON(http.StatusOK, gin.H{"message": "done"})
}
