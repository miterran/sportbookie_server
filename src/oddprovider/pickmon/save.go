package pickmon

import (
	"context"
	"log"
	"math"
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/model"
	"sport_bookie_server/src/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Save ...
func Save(lines Lines) error {
	for _, game := range lines.Games {
		var homeTeam = model.TeamDetail{
			Rot:  game.HomeTeam.Rot,
			Name: game.HomeTeam.Name.Parse,
		}
		var awayTeam = model.TeamDetail{
			Rot:  game.AwayTeam.Rot,
			Name: game.AwayTeam.Name.Parse,
		}
		var team = model.Team{
			Home: homeTeam,
			Away: awayTeam,
		}
		var money = model.Money{
			HomeOdd: game.Line.Money.Home,
			AwayOdd: game.Line.Money.Away,
		}
		var draw = model.Draw{
			Odd: game.Line.Money.Draw,
		}
		var spread model.Spread
		if util.IsValidPoints(game.Line.Spread.Points) {
			var awayPoint float64
			if game.Line.Spread.Points > 0 {
				awayPoint = -game.Line.Spread.Points
			} else {
				awayPoint = math.Abs(game.Line.Spread.Points)
			}
			spread = model.Spread{
				HomePoints: game.Line.Spread.Points,
				AwayPoints: awayPoint,
				HomeOdd:    game.Line.Spread.Home,
				AwayOdd:    game.Line.Spread.Away,
			}
		}
		var total model.Total
		if util.IsValidPoints(game.Line.Total.Points) {
			total = model.Total{
				Points:   game.Line.Total.Points,
				OverOdd:  game.Line.Total.Over,
				UnderOdd: game.Line.Total.Under,
			}
		}
		var line = model.Line{
			Money:  money,
			Spread: spread,
			Total:  total,
			Draw:   draw,
		}
		var score = model.Score{
			Home: game.Line.Score.Home,
			Away: game.Line.Score.Away,
		}
		var status = 0
		if game.Line.Score.Winner != "" {
			status = 1
		}
		if game.Void != 0 {
			status = 2
		}
		var newGame = model.Game{
			Provider:    "pickmon",
			ProviderID:  game.ID,
			Sport:       game.Sport.Parse,
			League:      game.League.Parse,
			MatchTime:   game.MatchTime.Parse,
			Team:        team,
			Period:      game.Line.Periodnum,
			CutOffTime:  game.Line.CutOffTime.Parse,
			Line:        line,
			Score:       score,
			Status:      status,
			LastUpdated: time.Now(),
		}
		if game.Sport.Parse == "" || game.League.Parse == "" {
			continue
		}
		filter := bson.M{"provider": "pickmon", "providerID": newGame.ProviderID, "status": 0}
		opts := options.FindOneAndReplace()
		opts.SetUpsert(true)
		db.Games.FindOneAndReplace(context.TODO(), filter, newGame, opts)
		log.Printf("update %v, %v, %v, %v, %v\n", "Pickmon", newGame.ProviderID, newGame.Sport, newGame.League, newGame.Status)
	}
	return nil
}
