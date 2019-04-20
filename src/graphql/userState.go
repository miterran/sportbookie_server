package graphql

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/model"
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

		state, openBetsWithGame, err := db.GetUserState(params.Context, userID)
		if err != nil {
			return nil, err
		}

		var user model.User
		err = db.Users.FindOne(params.Context, bson.M{"_id": userID}).Decode(&user)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &model.UserState{
			Initial:          user.InitialCredit,
			Balance:          state["balance"],
			AtRisk:           state["atRisk"],
			ToWin:            state["toWin"],
			Available:        user.InitialCredit + state["balance"] - state["atRisk"],
			OpenBetsWithGame: openBetsWithGame,
		}, nil
	},
}
