package handler

import (
	"net/http"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/gin-gonic/gin"
)

func CreateRequisition(c *gin.Context) {
	var requisition model.RequisitionDTO
	if err := c.BindJSON(&requisition); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}
