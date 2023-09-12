package handler

import (
	"net/http"
	"strconv"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/gin-gonic/gin"
)

var productService = &service.ProductService{}

func AddProduct(c *gin.Context) {
	var products []model.Product

	if err := c.BindJSON(&products); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := productService.AddProducts(products); err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	c.String(http.StatusCreated, "Products added!")
}

func SearchProducts(c *gin.Context) {
	searchString := c.Query("searchString")

	products, err := productService.SearchForProducts(searchString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProductPrice(c *gin.Context) {
	var price model.ProductDTO

	if err := c.BindJSON(&price); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := productService.UpdateProductPrice(price); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Price updated",
	})
}

func GetProductPriceStats(c *gin.Context) {
	productIdString := c.Query("id")
	productId, err := strconv.ParseInt(productIdString, 10, 64)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	prices, err := productService.GetProductPriceStats(productId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, prices)
}
