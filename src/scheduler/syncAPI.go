package scheduler

import (
	"fmt"
	"log"
	"sport_bookie_server/src/config"
	"sport_bookie_server/src/oddprovider/pickmon"
	"sport_bookie_server/src/result"
	"time"
)

var gameURL = fmt.Sprintf("https://api.pickmonitor.com/lines.php?uid=%v&key=%v&graded=%v&full_call=1", config.PICKMONUID, config.PICKMONKEY, "0")
var scoreURL = fmt.Sprintf("https://api.pickmonitor.com/lines.php?uid=%v&key=%v&graded=%v&full_call=1", config.PICKMONUID, config.PICKMONKEY, "1")

// UpdateGame ...
func UpdateGame() {
	lines, err := pickmon.FetchGames(gameURL)
	if err != nil {
		log.Fatal(err)
	}
	err = pickmon.Save(lines)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateScore ...
func UpdateScore() {
	lines, err := pickmon.FetchGames(scoreURL)
	if err != nil {
		log.Fatal(err)
	}
	err = pickmon.Save(lines)
	if err != nil {
		log.Fatal(err)
	}
}

// SyncProvider ...
func SyncProvider() {
	go func() {
		for range time.Tick(config.SYNCGAMEDURATION) {
			UpdateGame()
		}
	}()
	go func() {
		for range time.Tick(config.SYNCSCOREDURATION) {
			UpdateScore()
			result.FinalOpenBetResult()
		}
	}()
}
