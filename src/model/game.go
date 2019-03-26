package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"log"
)

// Score ...
type Score struct {
	Home float64 `json:"home" bson:"home"`
	Away float64 `json:"away" bson:"away"`
}

// Money ...
type Money struct {
	HomeOdd float64 `json:"homeOdd" bson:"homeOdd"`
	AwayOdd float64 `json:"awayOdd" bson:"awayOdd"`
}

// Draw ...
type Draw struct {
	Odd float64 `json:"odd" bson:"odd"`
}

// Spread ...
type Spread struct {
	HomePoints float64 `json:"homePoints" bson:"homePoints"`
	AwayPoints float64 `json:"awayPoints" bson:"awayPoints"`
	HomeOdd float64 `json:"homeOdd" bson:"homeOdd"`
	AwayOdd float64 `json:"awayOdd" bson:"awayOdd"`
}

// Total ...
type Total struct {
	Points float64 `json:"points" bson:"points"`
	OverOdd float64 `json:"overOdd" bson:"overOdd"`
	UnderOdd float64 `json:"underOdd" bson:"underOdd"`
}

// Line ...
type Line struct {
	Money Money `json:"money" bson:"money"`
	Spread Spread `json:"spread" bson:"spread"`
	Total Total `json:"total" bson:"total"`
	Draw Draw `json:"draw" bson:"draw"`
}

// TeamDetail ...
type TeamDetail struct {
	Rot string `json:"rot" bson:"rot"`
	Name string `json:"name" bson:"name"`
}

// Team ...
type Team struct {
	Home TeamDetail `json:"home" bson:"home"`
	Away TeamDetail `json:"away" bson:"away"`
}

// Game ...
type Game struct {
	ID           primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Provider string `json:"provider" bson:"provider"`
	ProviderID string `json:"providerID" bson:"providerID"`
	Sport  string   `json:"sport" bson:"sport"`
	League string `json:"league" bson:"league"`
	MatchTime time.Time `json:"matchTime" bson:"matchTime"`
	Team Team `json:"team" bson:"team"`
	Period int `json:"period" bson:"period"`
	CutOffTime time.Time `json:"cutOffTime" bson:"cutOffTime"`
	Line Line `json:"line" bson:"line"`
	Score Score `json:"score" bson:"score"`
	Status int `json:"status" bson:"status"`
	LastUpdated time.Time `json:"lastUpdated" bson:"lastUpdated"`
}

// CompareLatestLine ...
func (l *Line) CompareLatestLine(selectedLineType string, selectedPointsType string, selectedPoints float64, selectOddType string, selectedOdd float64) bool {
	latestLine := map[string]map[string]float64 {
		"money": map[string]float64{
			"homeOdd": l.Money.HomeOdd,
			"awayOdd": l.Money.AwayOdd,
		},
		"spread": map[string]float64{
			"homePoints": l.Spread.HomePoints,
			"awayPoints": l.Spread.AwayPoints,
			"homeOdd": l.Spread.HomeOdd,
			"awayOdd": l.Spread.AwayOdd,
		},
		"total": map[string]float64{
			"points": l.Total.Points,
			"overOdd": l.Total.OverOdd,
			"underOdd": l.Total.UnderOdd,
		},
		"draw": map[string]float64{
			"odd": l.Draw.Odd,
		},
	}
	log.Println(latestLine[selectedLineType][selectOddType], selectedOdd, latestLine[selectedLineType][selectedPointsType], selectedPoints)
	return latestLine[selectedLineType][selectOddType] == selectedOdd && latestLine[selectedLineType][selectedPointsType] == selectedPoints
}

