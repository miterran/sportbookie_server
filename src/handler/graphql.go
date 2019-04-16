package handler

import (
	"sport_bookie_server/src/graphql"
	"sport_bookie_server/src/jwt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// Graphqlhandler handle graphql
func Graphqlhandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		c.Set("ID", claims["ID"].(string))
		h.ContextHandler(c, c.Writer, c.Request)
	}
}
