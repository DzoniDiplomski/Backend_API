package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var requisitionService = service.RequisitionService{}

func CreateRequisition(c *gin.Context) {
	var requisition model.RequisitionDTO
	if err := c.BindJSON(&requisition); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	jsonval, _ := json.Marshal(requisition)
	os.WriteFile("gas.json", jsonval, 0666)

	if err := requisitionService.CreateRequisition(requisition); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func GetRequisitions(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit
	requisitions, err := requisitionService.GetRequisitions(offset, limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, requisitions)
}

func CalculatePagesForAllRequisitions(c *gin.Context) {
	itemsPerPageStr := c.DefaultQuery("items", "10")
	itemsPerPage, _ := strconv.Atoi(itemsPerPageStr)

	pageStructure, err := requisitionService.CalculatePagesForAllRequisitions(itemsPerPage)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, pageStructure)
}

func GetRequisitionItems(c *gin.Context) {
	requisitionIdString := c.Query("id")
	requisitionId, _ := strconv.ParseInt(requisitionIdString, 10, 64)

	requisitionItems, err := requisitionService.GetRequisitionItems(requisitionId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, requisitionItems)
}
