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
	requisitions, err := requisitionService.GetRequisitions()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, requisitions)
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
