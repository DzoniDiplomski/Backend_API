package handler

import (
	"net/http"
	"strconv"

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

func GetReceipts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	receipts, err := receiptService.GetReceiptsWithLimit(offset, limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, receipts)
}

func CalculatePagesForAllReceipts(c *gin.Context) {
	itemsPerPageStr := c.DefaultQuery("items", "10")
	itemsPerPage, _ := strconv.Atoi(itemsPerPageStr)

	pageStructure, err := receiptService.CalculatePagesForAllReceipts(itemsPerPage)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, pageStructure)
}
