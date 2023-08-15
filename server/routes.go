package server

import (
	"github.com/DzoniDiplomski/Backend_API/handler"
	"github.com/gin-gonic/gin"
)

func initAuthHandlerRoutes(r *gin.Engine) {
	r.POST("/login", handler.Login)
}

func initProductHandlerRoutes(r *gin.Engine) {
	r.POST("/addProducts", handler.AddProduct)
	r.GET("/getProducts", handler.SearchProducts)
}

func initReceiptHandlerRoutes(r *gin.Engine) {
	r.POST("/addReceipt", handler.CreateReceipt)
}
