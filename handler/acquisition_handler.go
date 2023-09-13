package handler

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var acquisitionService = service.AcquisitionService{}

func GetAcquisitions(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit
	acquisitionNames, err := acquisitionService.GetAcquisitionNames(offset, limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, acquisitionNames)
}

func CalculatePagesForAllAcquisitions(c *gin.Context) {
	itemsPerPageStr := c.DefaultQuery("items", "10")
	itemsPerPage, _ := strconv.Atoi(itemsPerPageStr)

	pageStructure, err := acquisitionService.CalculatePagesForAllAcquisitions(itemsPerPage)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, pageStructure)
}

func OpenAcquisition(c *gin.Context) {
	filename := c.Query("filename")
	filename = strings.ReplaceAll(filename, "%", " ")
	pdfData, err := acquisitionService.OpenAcquisition(filename)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	pdfBase64 := base64.StdEncoding.EncodeToString(pdfData)
	responseData := map[string]string{"pdfData": pdfBase64}

	c.JSON(http.StatusOK, responseData)
}
