package graphql

import (
	"sport_bookie_server/src/model"
	"github.com/graphql-go/graphql"
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// WagerType ...
var WagerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WagerType",
		Fields: graphql.Fields{
			"toWin": &graphql.Field{
				Type: graphql.Int,
			},
			"atRisk": &graphql.Field{
				Type: graphql.Int,
			},

		},
	},
)

// SelectedType ...
var SelectedType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SelectedType",
		Fields: graphql.Fields{
			"lineType": &graphql.Field{
				Type: graphql.String,
			},
			"pointsType": &graphql.Field{
				Type: graphql.String,
			},
			"points": &graphql.Field{
				Type: graphql.Float,
			},
			"oddType": &graphql.Field{
				Type: graphql.String,
			},
			"odd": &graphql.Field{
				Type: graphql.Float,
			},
			"target": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// BetWithGameType ...
var BetWithGameType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BetWithGameType",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: OIDScalarType,
			},
			"userID": &graphql.Field{
				Type: OIDScalarType,
			},
			"gameID": &graphql.Field{
				Type: OIDScalarType,
			},
			"game": &graphql.Field{
				Type: GameType,
			},
			"selected": &graphql.Field{
				Type: SelectedType,
			},
			"wager": &graphql.Field{
				Type: WagerType,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"balance": &graphql.Field{
				Type: graphql.Int,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"lastUpdated": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// HistoryBetsType ...
var HistoryBetsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "HistoryBetsType",
		Fields: graphql.Fields{
			"fromDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"toDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"balance": &graphql.Field{
				Type: graphql.Int,
			},
			"betsWithGame": &graphql.Field{
				Type: graphql.NewList(BetWithGameType),
			},
		},
	},
)

// HistorysBetsQuery ...
var HistorysBetsQuery = &graphql.Field{
	Type:        graphql.NewList(HistoryBetsType),
	Description: "get user's historys bets",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		uid := params.Context.Value("ID").(string)
		userID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			return nil, err
		}
		var historysBets []model.HistoryBets
		for i := 0; i < 8; i++ {
			var balance int
			year, week := time.Now().AddDate(0,0,-(i*7)).ISOWeek()
			fromDate, toDate := util.ISOWeekRange(year, week)
			betsWithGame, err := db.GetUserHistoryBetsFromISOWeek(params.Context, userID, year, week)
			if err != nil {
				return nil, err
			}
			for _, betWithGame := range betsWithGame {
				balance += betWithGame.Balance
			}
			var historyBets = model.HistoryBets {
				FromDate: fromDate,
				ToDate: toDate,
				Balance: balance,
				BetsWithGame: betsWithGame,
			}
			historysBets = append(historysBets, historyBets)
		}
		return historysBets, nil
	},
}