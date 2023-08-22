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
	r.PUT("/updatePrice", handler.UpdateProductPrice)
}

func initReceiptHandlerRoutes(r *gin.Engine) {
	r.POST("/addReceipt", handler.CreateReceipt)
	r.GET("/getAllReceipts", handler.GetReceipts)
	r.GET("/calculatePages", handler.CalculatePagesForAllReceipts)
	r.GET("/getAllInvoices", handler.GetInvoices)
	r.GET("/calculateInvoicePages", handler.CalculatePagesForAllInvoices)
}

func initRequisitionHandlerRoutes(r *gin.Engine) {
	r.POST("/addRequisition", handler.CreateRequisition)
}
