package handler

import (
	"net/http"

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

	if err := requisitionService.CreateRequisition(requisition); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
