package graphql

import (
	"log"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/graphql-go/graphql/language/ast"
)

// OIDScalarType ...
var OIDScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "OIDScalarType",
	Description: "Conver Object ID to String.",
	// Serialize serializes `CustomID` to string.
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case primitive.ObjectID:
			return value.Hex()
		case *primitive.ObjectID:
			return value.Hex()
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `CustomID`.
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			oid, err := primitive.ObjectIDFromHex(value)
			if err != nil {
				return nil
			}
			return oid
		case *string:
			oid, err := primitive.ObjectIDFromHex(*value)
			if err != nil {
				return nil
			}
			return oid
		default:
			return nil
		}
	},
	// ParseLiteral parses GraphQL AST value to `CustomID`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		log.Println("graphql OIDScalar type")
		log.Println(valueAST)
		return nil
	},
})