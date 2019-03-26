package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID           primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username,omitempty" form:"username" binding:"required"`
	Password     string             `json:"password" bson:"password,omitempty" form:"password" binding:"required"`
	InitialCredit int 				`json:"initialCredit" bson:"initialCredit"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt" `
	LastOnlineAt time.Time          `json:"lastOnlineAt" bson:"lastOnlineAt"`
}

// UserState ...
type UserState struct {
	Initial   int `json:"initial"`
	Balance   int `json:"balance"`
	AtRisk  int `json:"atRisk"`
	ToWin int `json:"toWin"`
	Available int `json:"available"`
	OpenBetsWithGame []BetWithGame `json:"betsWithGame"`
}