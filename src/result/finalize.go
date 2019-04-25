package result

import (
	"context"
	"log"
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// Finalize Open Bet Result ...
func FinalizeResult() {
	betsWithGame, _ := db.FindBets(context.TODO(), bson.M{"status": 0})
	var wg sync.WaitGroup
	wg.Add(len(betsWithGame))
	for _, betWithGame := range betsWithGame {
		go findOneAndReplace(&wg, betWithGame)
	}
	wg.Wait()
}

func findOneAndReplace(wg *sync.WaitGroup, betWithGame model.BetWithGame) {
	defer wg.Done()
	if betWithGame.Game.Status == 0 {
		return
	}
	var newBetBalance = 0
	var newBetStatus = 0
	if betWithGame.Game.Status == 1 {
		homeScore := betWithGame.Game.Score.Home
		awayScore := betWithGame.Game.Score.Away
		totalScore := homeScore + awayScore
		selectedLineType := betWithGame.Selected.LineType
		selectedPointsType := betWithGame.Selected.PointsType
		selectedOddType := betWithGame.Selected.OddType
		selectedPoints := betWithGame.Selected.Points
		var spreadPointsAdjust = map[string]float64{
			"homePoints": homeScore + selectedPoints,
			"awayPoints": awayScore + selectedPoints,
		}
		var won = map[string]map[string]bool{
			"money": {
				"homeOdd": homeScore > awayScore,
				"awayOdd": awayScore > homeScore,
			},
			"spread": {
				"homeOdd": spreadPointsAdjust[selectedPointsType] > awayScore,
				"awayOdd": spreadPointsAdjust[selectedPointsType] > homeScore,
			},
			"total": {
				"overOdd":  totalScore > selectedPoints,
				"underOdd": totalScore < selectedPoints,
			},
			"draw": {
				"odd": homeScore == awayScore,
			},
		}
		if won[selectedLineType][selectedOddType] {
			newBetBalance = betWithGame.Wager.ToWin
		} else {
			newBetBalance = -betWithGame.Wager.AtRisk
		}
		newBetStatus = 1
	}
	if betWithGame.Game.Status == 2 {
		newBetStatus = 2
	}
	var newBet = model.Bet{
		UserID:      betWithGame.UserID,
		GameID:      betWithGame.GameID,
		Selected:    betWithGame.Selected,
		Wager:       betWithGame.Wager,
		Status:      newBetStatus,
		Balance:     newBetBalance,
		CreatedAt:   betWithGame.CreatedAt,
		LastUpdated: time.Now(),
	}
	db.Bets.FindOneAndReplace(context.TODO(), bson.M{"_id": betWithGame.ID}, newBet)
	log.Printf("update %v, %v, %v\n", "result", betWithGame.UserID, newBet.Balance)

}
