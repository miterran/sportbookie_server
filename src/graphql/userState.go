package graphql

import (
	"sport_bookie_server/src/model"
	"github.com/graphql-go/graphql"
	"log"
	"sport_bookie_server/src/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// UserStateType ...
var UserStateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserStateType",
		Fields: graphql.Fields{
			"initial": &graphql.Field{
				Type: graphql.Int,
			},
			"toWin": &graphql.Field{
				Type: graphql.Int,
			},
			"atRisk": &graphql.Field{
				Type: graphql.Int,
			},
			"balance": &graphql.Field{
				Type: graphql.Int,
			},
			"available": &graphql.Field{
				Type: graphql.Int,
			},
			"betsWithGame": &graphql.Field{
				Type: graphql.NewList(BetWithGameType),
			},
		},
	},
)

// UserStateQuery ...
var UserStateQuery = &graphql.Field{
	Type:        UserStateType,
	Description: "get user's state",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		uid := params.Context.Value("ID").(string)
		userID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			return nil, err
		}
		var state = map[string]int {
			"toWin": 0,
			"atRisk": 0,
			"balance": 0,
		}
		openBetsWithGame, err := db.GetUserOpenBets(params.Context, userID)
		if err != nil {
			return nil, err
		}
		for _, openBetWithGame := range openBetsWithGame {
			state["toWin"] += openBetWithGame.Wager.ToWin
			state["atRisk"] += openBetWithGame.Wager.AtRisk
		}
		year, week := time.Now().ISOWeek()
		currentWeekHistoryBetsWithGame, err := db.GetUserHistoryBetsFromISOWeek(params.Context, userID, year, week)
		if err != nil {
			return nil, err
		}
		for _, currentWeekHistoryBetWithGame := range currentWeekHistoryBetsWithGame {
			state["balance"] += currentWeekHistoryBetWithGame.Balance
		}
		var user model.User
		err = db.Users.FindOne(params.Context, bson.M{"_id": userID}).Decode(&user)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &model.UserState{
			Initial:   user.InitialCredit,
			Balance:   state["balance"],
			AtRisk:   state["atRisk"],
			ToWin:   state["toWin"],
			Available: user.InitialCredit + state["balance"] - state["atRisk"],
			OpenBetsWithGame: openBetsWithGame,
		}, nil
	},
}


