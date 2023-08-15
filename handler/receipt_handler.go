package handler

import (
	"net/http"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var receiptService = &service.ReceiptService{}

func CreateReceipt(c *gin.Context) {
	var receipt model.ReceiptDTO
	if err := c.BindJSON(&receipt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := receiptService.CreateReceipt(receipt); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
