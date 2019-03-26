package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Client ...
var Client *mongo.Client

// Users = UsersCollection
var Users *mongo.Collection

// Picks = PicksCollection
var Picks *mongo.Collection

// Bets = BetsCollection
var Bets *mongo.Collection

// Games = EventsCollection
var Games *mongo.Collection

// Connect to mongodb server and setup collections
func Connect(URI string) {
	opt := options.Client().ApplyURI(URI)
	connection, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		log.Fatal(err)
	}
	err = connection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	Client = connection
	Users = Client.Database("sportbookie").Collection("users")
	Picks = Client.Database("sportbookie").Collection("picks")
	Games = Client.Database("sportbookie").Collection("games")
	Bets = Client.Database("sportbookie").Collection("bets")
	indexesSetup()
	log.Println("Connected to MongoDB!")
}

// indexesSetup set Users username unique
func indexesSetup() {
	_, err := Users.Indexes().DropAll(context.TODO())
	if err != nil {
		log.Println(err)
	}
	indexOption := options.Index()
	indexOption.SetUnique(true)
	indexOption.SetName("username")
	indexModel := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "username", Value: bsonx.Int32(1)}},
		Options: indexOption,
	}
	_, err = Users.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Println(err)
	}
}