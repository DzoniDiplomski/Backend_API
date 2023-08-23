package service

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type DailyMarketService struct {
}

func (dailyMarketService *DailyMarketService) GetDailyMarketService() (model.DailyMarket, error) {
	var dailyMarket model.DailyMarket
	err := db.DBConn.QueryRow(db.PSGetTodaysMarket).Scan(&dailyMarket.Id, &dailyMarket.Date, &dailyMarket.Sum)
	if err != nil {
		return model.DailyMarket{}, err
	}
	return dailyMarket, nil
}
