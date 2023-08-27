package handler

import (
	"net/http"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var calculationService = service.CalculationService{}

func CreateCalculation(c *gin.Context) {
	var calculation model.Calculation
	if err := c.BindJSON(&calculation); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := calculationService.CreateCalculation(calculation); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusOK)
}
