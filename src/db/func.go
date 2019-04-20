package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sport_bookie_server/src/model"
	"sport_bookie_server/src/util"
	"time"
)

// GetUserOpenBets ...
func GetUserOpenBets(c context.Context, userID primitive.ObjectID) ([]model.BetWithGame, error) {
	filter := bson.M{"userID": userID, "status": 0}
	betsWithGame, err := FindBets(c, filter)
	if err != nil {
		return nil, err
	}
	return betsWithGame, nil
}

// GetUserHistoryBetsFromISOWeek ...
func GetUserHistoryBetsFromISOWeek(c context.Context, userID primitive.ObjectID, year int, week int) ([]model.BetWithGame, error) {
	fromDate, toDate := util.ISOWeekRange(year, week)
	filter := bson.M{
		"userID": userID,
		"createdAt": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
		"status": bson.M{
			"$ne": 0,
		},
	}
	betsWithGame, err := FindBets(c, filter)
	if err != nil {
		return nil, err
	}
	return betsWithGame, nil
}

// FindBets ...
func FindBets(c context.Context, filter bson.M) ([]model.BetWithGame, error) {
	cur, err := Bets.Find(c, filter)
	defer cur.Close(c)
	if err != nil {
		return nil, err
	}

	var betsWithGame []model.BetWithGame
	for cur.Next(c) {
		var bet model.Bet
		err := cur.Decode(&bet)
		if err != nil {
			return nil, err
		}
		var game model.Game
		err = Games.FindOne(c, bson.M{"_id": bet.GameID}).Decode(&game)

		if err != nil {
			return nil, err
		}
		var betWithGame = model.BetWithGame{
			ID:          bet.ID,
			UserID:      bet.UserID,
			GameID:      bet.GameID,
			Selected:    bet.Selected,
			Wager:       bet.Wager,
			Status:      bet.Status,
			Balance:     bet.Balance,
			CreatedAt:   bet.CreatedAt,
			LastUpdated: bet.LastUpdated,
			Game:        game,
		}
		betsWithGame = append([]model.BetWithGame{betWithGame}, betsWithGame...)
	}
	return betsWithGame, nil
}

// GetUserState get user current week state .. toWin | atRisk | balance
func GetUserState(c context.Context, userID primitive.ObjectID) (map[string]int, []model.BetWithGame, error) {
	var state = map[string]int{
		"toWin":   0,
		"atRisk":  0,
		"balance": 0,
	}
	openBetsWithGame, err := GetUserOpenBets(c, userID)
	if err != nil {
		return nil, nil, err
	}
	for _, openBetWithGame := range openBetsWithGame {
		state["toWin"] += openBetWithGame.Wager.ToWin
		state["atRisk"] += openBetWithGame.Wager.AtRisk
	}
	year, week := time.Now().ISOWeek()
	currentWeekHistoryBetsWithGame, err := GetUserHistoryBetsFromISOWeek(c, userID, year, week)
	if err != nil {
		return nil, nil, err
	}
	for _, currentWeekHistoryBetWithGame := range currentWeekHistoryBetsWithGame {
		state["balance"] += currentWeekHistoryBetWithGame.Balance
	}
	return state, openBetsWithGame, nil
}
