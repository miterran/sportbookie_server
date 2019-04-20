package config

import (
	"log"
	"os"
	"time"
)

// MONGOURL env
var MONGOURL string

// PORT env
var PORT string

// PICKMONUID env
var PICKMONUID string

// PICKMONKEY env
var PICKMONKEY string

// SYNCCD env
var SYNCCD time.Duration

func init() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("missing env PORT, ie: 8080")
	}

	MONGOURL = os.Getenv("MONGO_URL")
	if MONGOURL == "" {
		log.Fatal("missing env PORT, ie: mongodb://username:password@ds000000.mlab.com:00000/db")
	}

	PICKMONUID = os.Getenv("PICKMON_UID")
	if PICKMONUID == "" {
		log.Fatal("missing env PICKMONUID, API UID from www.pickmonitor.com")
	}

	PICKMONKEY = os.Getenv("PICKMON_KEY")
	if PICKMONUID == "" {
		log.Fatal("missing env PICKMONKEY, API KEY from www.pickmonitor.com")
	}

	synccd := os.Getenv("SYNC_CD")
	if synccd == "" {
		synccd = "5m"
	}

	synccdduration, err := time.ParseDuration(synccd)
	if err != nil {
		log.Fatal(err)
	}
	if synccdduration.Minutes() < 3 {
		log.Fatal("SYNC COOL DOWN must greater then 3 minutes")
	}
	if synccdduration.Minutes() > 60 {
		log.Fatal("SYNC COOL DOWN must lesser then 60 minutes")
	}
	SYNCCD = synccdduration

}
