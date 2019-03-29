package graphql

import (
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/model"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BetOrderStatusType ...
var BetOrderStatusType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BetOrderStatusType",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// BetOrderStatus ...
type BetOrderStatus struct {
	Code int `josn:"code"`
}

// SubmitBetOrder ...
var SubmitBetOrder = &graphql.Field{
	Type:        BetOrderStatusType,
	Description: "Create new bet order",
	Args: graphql.FieldConfigArgument{
		"selectedGameID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"selectedLineType": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"selectedPointsType": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"selectedPoints": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"selectOddType": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"selectedOdd": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"selectedTarget": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"atRisk": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"toWin": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		atRisk, _ := params.Args["atRisk"].(int)
		if atRisk < 10 {
			return &BetOrderStatus{2}, nil
		}
		uid := params.Context.Value("ID").(string)
		userID, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			return nil, err
		}
		selectedGameID, _ := params.Args["selectedGameID"].(string)
		selectedLineType, _ := params.Args["selectedLineType"].(string)
		selectedPointsType, _ := params.Args["selectedPointsType"].(string)
		selectedPoints, _ := params.Args["selectedPoints"].(float64)
		selectOddType, _ := params.Args["selectOddType"].(string)
		selectedOdd, _ := params.Args["selectedOdd"].(float64)
		selectedTarget, _ := params.Args["selectedTarget"].(string)
		toWin, _ := params.Args["toWin"].(int)
		var game model.Game
		gameID, err := primitive.ObjectIDFromHex(selectedGameID)
		if err != nil {
			return nil, err
		}
		err = db.Games.FindOne(params.Context, bson.M{"_id": gameID}).Decode(&game)
		if err != nil {
			return &BetOrderStatus{6}, nil
		}
		if game.CutOffTime.Before(time.Now()) || game.Status != 0 {
			return &BetOrderStatus{5}, nil
		}
		if !game.Line.CompareLatestLine(selectedLineType, selectedPointsType, selectedPoints, selectOddType, selectedOdd) {
			return &BetOrderStatus{4}, nil
		}
		selected := model.Selected{
			LineType:   selectedLineType,
			PointsType: selectedPointsType,
			Points:     selectedPoints,
			OddType:    selectOddType,
			Odd:        selectedOdd,
			Target:     selectedTarget,
		}
		wager := model.Wager{
			AtRisk: atRisk,
			ToWin:  toWin,
		}
		newBet := model.Bet{
			UserID:      userID,
			GameID:      gameID,
			Selected:    selected,
			Wager:       wager,
			Status:      0,
			Balance:     0,
			CreatedAt:   time.Now(),
			LastUpdated: time.Now(),
		}
		_, err = db.Bets.InsertOne(params.Context, newBet)
		if err != nil {
			return nil, err
		}
		return &BetOrderStatus{1}, nil
	},
}

// 0 default
// 1 success saved, reidrect to home
// 2 min risk
// 3 not enought credit
// 4 odd updated
// 5 timeout
// 6 game not found
// else server error
