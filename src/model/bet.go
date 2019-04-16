package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Selected ...
type Selected struct {
	LineType   string  `json:"lineType" bson:"lineType"`
	PointsType string  `json:"pointsType" bson:"pointsType"`
	Points     float64 `json:"points" bson:"points"`
	OddType    string  `json:"oddType" bson:"oddType"`
	Odd        float64 `json:"odd" bson:"odd"`
	Target     string  `json:"target" bson:"target"`
}

// Wager ...
type Wager struct {
	ToWin  int `json:"toWin" bson:"toWin"`
	AtRisk int `json:"atRisk" bson:"atRisk"`
}

// Bet ...
type Bet struct {
	ID          primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userID" bson:"userID,omitempty"`
	GameID      primitive.ObjectID `json:"gameID" bson:"gameID,omitempty"`
	Selected    Selected           `json:"selected" bson:"selected"`
	Wager       Wager              `json:"wager" bson:"wager"`
	Status      int                `json:"status" bson:"status"`
	Balance     int                `json:"balance" bson:"balance"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	LastUpdated time.Time          `json:"lastUpdated" bson:"lastUpdated"`
}

// BetWithGame ...
type BetWithGame struct {
	ID          primitive.ObjectID `json:"ID"`
	UserID      primitive.ObjectID `json:"userID"`
	GameID      primitive.ObjectID `json:"gameID"`
	Selected    Selected           `json:"selected"`
	Wager       Wager              `json:"wager"`
	Status      int                `json:"status"`
	Balance     int                `json:"balance"`
	CreatedAt   time.Time          `json:"createdAt"`
	LastUpdated time.Time          `json:"lastUpdated"`
	Game        Game               `json:"game"`
}

// OpenBets model
type OpenBets struct {
	AtRisk       int           `json:"toWin"`
	ToWin        int           `json:"atRisk"`
	BetsWithGame []BetWithGame `json:"betsWithGame"`
}

// HistoryBets model
type HistoryBets struct {
	FromDate     time.Time     `json:"fromDate"`
	ToDate       time.Time     `json:"toDate"`
	Balance      int           `json:"balance"`
	BetsWithGame []BetWithGame `json:"betsWithGame"`
}

// status
// 0 pending
// 1 done
// 2 cancelled
// 3 error
