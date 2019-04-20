package middleware

import (
	"sport_bookie_server/src/oddprovider/pickmon"
	"github.com/gin-gonic/gin"
)

// SyncProviderData ...
func SyncProviderData() gin.HandlerFunc {
	return func(c *gin.Context) {
		pickmon.Sync()
		c.Next()
	}
}