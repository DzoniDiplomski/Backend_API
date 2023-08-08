package handler

import (
	"net/http"

	"diplomski.com/model"
	"diplomski.com/service"
	"github.com/gin-gonic/gin"
)

var productService = &service.ProductService{}

func AddProduct(c *gin.Context) {
	var products []model.Product

	if err := c.BindJSON(&products); err != nil {
		c.Status(400)
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
