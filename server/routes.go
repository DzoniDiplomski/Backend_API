package server

import (
	"github.com/DzoniDiplomski/Backend_API/handler"
	"github.com/DzoniDiplomski/Backend_API/middleware"
	"github.com/gin-gonic/gin"
)

func initAuthHandlerRoutes(r *gin.Engine) {
	r.POST("/login", handler.Login)
}

func initCashierRoutes(r *gin.Engine) {
	cashierGroup := r.Group("/cashier")
	cashierGroup.Use(middleware.CashierMiddleware)
	cashierGroup.GET("/getProducts", handler.SearchProducts)
	cashierGroup.POST("/addReceipt", handler.CreateReceipt)
	cashierGroup.GET("/getAllReceipts", handler.GetReceipts)
	cashierGroup.GET("/calculatePages", handler.CalculatePagesForAllReceipts)
	cashierGroup.GET("/getAllInvoices", handler.GetInvoices)
	cashierGroup.GET("/calculateInvoicePages", handler.CalculatePagesForAllInvoices)
}

func initManagerRoutes(r *gin.Engine) {
	managerGroup := r.Group("/manager")
	// managerGroup.Use(middleware.ManagerMiddleware)
	managerGroup.POST("/addRequisition", handler.CreateRequisition)
	managerGroup.PUT("/updatePrice", handler.UpdateProductPrice)
	managerGroup.GET("/priceStats", handler.GetProductPriceStats)
	managerGroup.GET("/getRequisitions", handler.GetRequisitions)
	managerGroup.GET("/getRequisitionItems", handler.GetRequisitionItems)
	managerGroup.GET("/getAcquisitions", handler.GetAcquisitions)
	managerGroup.GET("/openAcquisition", handler.OpenAcquisition)
}

func initStorageWorkerRoutes(r *gin.Engine) {
	storageWorkerGroup := r.Group("/storageWorker")
	storageWorkerGroup.Use(middleware.StorageWorkerMiddleware)
	storageWorkerGroup.POST("/addProducts")
}

func initDailyMarketRoutes(r *gin.Engine) {
	r.GET("/getDailyMarket", handler.GetDailyMarket)
}
