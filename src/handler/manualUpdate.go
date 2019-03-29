package handler

import (
	"log"
	"net/http"
	"sport_bookie_server/src/result"
	"sport_bookie_server/src/scheduler"

	"github.com/gin-gonic/gin"
)

// ManualUpdate manual update db
func ManualUpdate(c *gin.Context) {
	scheduler.UpdateGame()
	scheduler.UpdateScore()
	result.FinalOpenBetResult()
	log.Println("done")
	c.JSON(http.StatusOK, gin.H{"message": "done"})
}
