package main

import (
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/router"
	"sport_bookie_server/src/config"
)

func main() {
	db.Connect(config.MONGOURL)
	router.Run(":"+config.PORT)
}
