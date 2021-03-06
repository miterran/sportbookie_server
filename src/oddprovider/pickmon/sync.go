package pickmon

import (
	"fmt"
	"log"
	"sport_bookie_server/src/config"
	"sport_bookie_server/src/result"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)
var lastUpdate = time.Now().Add(-1 * config.SYNCCD)

var gameURL = fmt.Sprintf("https://api.pickmonitor.com/lines.php?uid=%v&key=%v&graded=%v&full_call=1", config.PICKMONUID, config.PICKMONKEY, "0")
var scoreURL = fmt.Sprintf("https://api.pickmonitor.com/lines.php?uid=%v&key=%v&graded=%v&full_call=1", config.PICKMONUID, config.PICKMONKEY, "1")

// UpdateGame ...
func UpdateGame() {
	lines, err := FetchGames(gameURL)
	if err != nil {
		log.Fatal(err)
	}
	Save(lines)
}

// UpdateScore ...
func UpdateScore() {
	lines, err := FetchGames(scoreURL)
	if err != nil {
		log.Fatal(err)
	}
	Save(lines)
}

// Sync ...
func Sync() {
	now := time.Now()
	if now.After(lastUpdate.Add(config.SYNCCD)) {
		lastUpdate = time.Now()
		wg.Add(2)
		go func() {
			UpdateGame()
			wg.Done()
		}()
		go func() {
			UpdateScore()
			result.FinalizeResult()
			wg.Done()
		}()
	}
	wg.Wait()
}
