package graphql

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootQuery",
	Fields: graphql.Fields{
		"historysBets": HistorysBetsQuery,
		"games": GamesQuery,
		"game": GameQuery,
		"userState": UserStateQuery,

	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootMutation",
	Fields: graphql.Fields{
 		"submitBetOrder": SubmitBetOrder,
	},
})

// Schema ...
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
