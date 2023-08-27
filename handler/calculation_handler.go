package handler

import (
	"net/http"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/gin-gonic/gin"
)

func CreateCalculation(c *gin.Context) {
	var calculation model.Calculation
	if err := c.BindJSON(&calculation); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

}
