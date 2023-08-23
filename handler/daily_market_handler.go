package handler

import (
	"net/http"

	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var dailyMarketService = service.DailyMarketService{}

func GetDailyMarket(c *gin.Context) {
	dailyMarket, err := dailyMarketService.GetDailyMarketService()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, dailyMarket)
}
