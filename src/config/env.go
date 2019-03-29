package config

import (
	"log"
	"os"
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
	if PORT == "" {
		log.Fatal("missing env PORT, ie: 8080")
	}
	MONGOURL = os.Getenv("MONGO_URL")
	if MONGOURL == "" {
		log.Fatal("missing env PORT, ie: mongodb://username:password@ds000000.mlab.com:00000/db")
	}
	PICKMONUID = os.Getenv("PICKMON_UID")
	if PICKMONUID == "" {
		log.Fatal("missing env PICKMONUID, API UID from www.pickmonitor.com")
	}
	PICKMONKEY = os.Getenv("PICKMON_KEY")
	if PICKMONUID == "" {
		log.Fatal("missing env PICKMONKEY, API KEY from www.pickmonitor.com")
	}

	gameDuration := os.Getenv("SYNC_GAME_DURATION")
	if gameDuration == "" {
		gameDuration = "2h"
	}
	syncGameDuration, err := time.ParseDuration(gameDuration)
	if err != nil {
		log.Fatal(err)
	}
	if syncGameDuration.Minutes() < 4 {
		log.Fatal("SYNC_GAME_DURATION must greater then 4 minutes")
	}
	SYNCGAMEDURATION = syncGameDuration

	scoreDuration := os.Getenv("SYNC_SCORE_DURATION")
	if scoreDuration == "" {
		scoreDuration = "4h"
	}
	syncScoreDuration, err := time.ParseDuration(scoreDuration)
	if err != nil {
		log.Fatal(err)
	}
	if syncScoreDuration.Minutes() < 30 {
		log.Fatal("SYNC_SCORE_DURATION must greater then 30 minutes")
	}
	SYNCSCOREDURATION = syncScoreDuration

}
