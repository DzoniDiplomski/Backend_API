package handler

import (
	"net/http"

	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var acquisitionService = service.AcquisitionService{}

func GetAcquisitions(c *gin.Context) {
	acquisitionNames, err := acquisitionService.GetAcquisitionNames()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, acquisitionNames)
}

func OpenAcquisition(c *gin.Context) {
	filename := c.Query("filename")
	pdfData, err := acquisitionService.OpenAcquisition(filename)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Data(http.StatusOK, "application/pdf", pdfData)
}
