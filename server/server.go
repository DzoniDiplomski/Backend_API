package server

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	// r.Use(middleware.AuthMiddleware)
	initRoutes(r)
	r.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	initAuthHandlerRoutes(r)
	initProductHandlerRoutes(r)
	initReceiptHandlerRoutes(r)
	initRequisitionHandlerRoutes(r)
}
