package graphql

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/model"
	"time"
)

// TeamType ...
var TeamType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TeamType",
		Fields: graphql.Fields{
			"home": &graphql.Field{
				Type: TeamDetailType,
			},
			"away": &graphql.Field{
				Type: TeamDetailType,
			},
		},
	},
)

// TeamDetailType ...
var TeamDetailType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TeamDetailType",
		Fields: graphql.Fields{
			"rot": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// ScoreType ...
var ScoreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ScoreType",
		Fields: graphql.Fields{
			"home": &graphql.Field{
				Type: graphql.Float,
			},
			"away": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// MoneyType ...
var MoneyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MoneyType",
		Fields: graphql.Fields{
			"homeOdd": &graphql.Field{
				Type: graphql.Float,
			},
			"awayOdd": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// DrawType ...
var DrawType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DrawType",
		Fields: graphql.Fields{
			"odd": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// SpreadType ...
var SpreadType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SpreadType",
		Fields: graphql.Fields{
			"homePoints": &graphql.Field{
				Type: graphql.Float,
			},
			"awayPoints": &graphql.Field{
				Type: graphql.Float,
			},
			"homeOdd": &graphql.Field{
				Type: graphql.Float,
			},
			"awayOdd": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// TotalType ...
var TotalType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TotalType",
		Fields: graphql.Fields{
			"points": &graphql.Field{
				Type: graphql.Float,
			},
			"overOdd": &graphql.Field{
				Type: graphql.Float,
			},
			"underOdd": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// LineType ...
var LineType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LineType",
		Fields: graphql.Fields{
			"money": &graphql.Field{
				Type: MoneyType,
			},
			"spread": &graphql.Field{
				Type: SpreadType,
			},
			"total": &graphql.Field{
				Type: TotalType,
			},
			"draw": &graphql.Field{
				Type: DrawType,
			},
		},
	},
)

// GameType ...
var GameType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GameType",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: OIDScalarType,
			},
			"provider": &graphql.Field{
				Type: graphql.String,
			},
			"providerID": &graphql.Field{
				Type: graphql.String,
			},
			"sport": &graphql.Field{
				Type: graphql.String,
			},
			"league": &graphql.Field{
				Type: graphql.String,
			},
			"matchTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"team": &graphql.Field{
				Type: TeamType,
			},
			"period": &graphql.Field{
				Type: graphql.Int,
			},
			"cutOffTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"line": &graphql.Field{
				Type: LineType,
			},
			"score": &graphql.Field{
				Type: ScoreType,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"lastUpdated": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// GamesQuery ...
var GamesQuery = &graphql.Field{
	Type:        graphql.NewList(GameType),
	Description: "Get upcoming games",
	Args: graphql.FieldConfigArgument{
		"sport": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		sport := params.Args["sport"].(string)
		filter := bson.M{
			"sport":  sport,
			"status": 0,
			"cutOffTime": bson.M{
				"$gt": time.Now(),
			},
		}
		options := options.FindOptions{}
		options.Sort = bson.D{primitive.E{Key: "matchTime", Value: 1}, primitive.E{Key: "providerID", Value: 1}}
		cur, err := db.Games.Find(params.Context, filter, &options)
		defer cur.Close(params.Context)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		var games []model.Game
		for cur.Next(params.Context) {
			var game model.Game
			err := cur.Decode(&game)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			games = append(games, game)
		}
		return games, nil
	},
}

// GameQuery ...
var GameQuery = &graphql.Field{
	Type:        GameType,
	Description: "Get single game",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		gameID := params.Args["ID"].(string)
		id, err := primitive.ObjectIDFromHex(gameID)
		if err != nil {
			return nil, err
		}
		var result model.Game
		err = db.Games.FindOne(params.Context, bson.M{"_id": id}).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	},
}
