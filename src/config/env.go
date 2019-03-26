package config

import (
	"os"
	"log"
	"time"
)

// MONGOURL env
var MONGOURL string
// PORT env
var PORT string
// PICKMONUID env
var PICKMONUID string
// PICKMONKEY env
var PICKMONKEY string
// SYNCGAMEDURATION env
var SYNCGAMEDURATION time.Duration
// SYNCSCOREDURATION env
var SYNCSCOREDURATION time.Duration

func init() {
	PORT = os.Getenv("PORT")
	MONGOURL = os.Getenv("MONGO_URL")

	PICKMONUID = os.Getenv("PICKMON_UID")
	PICKMONKEY = os.Getenv("PICKMON_KEY")

	gameDuration := os.Getenv("SYNC_GAME_DURATION")
	syncGameDuration, err := time.ParseDuration(gameDuration)
	if err != nil {
		log.Fatal(err)
	}
	if syncGameDuration.Minutes() < 2 {
		log.Fatal("SYNC_GAME_DURATION must greater then 2 minutes")
	}
	SYNCGAMEDURATION = syncGameDuration

	scoreDuration := os.Getenv("SYNC_SCORE_DURATION")
	syncScoreDuration, err := time.ParseDuration(scoreDuration)
	if err != nil {
		log.Fatal(err)
	}
	if syncScoreDuration.Minutes() < 30 {
		log.Fatal("SYNC_SCORE_DURATION must greater then 30 minutes")
	}
	SYNCSCOREDURATION = syncScoreDuration

}