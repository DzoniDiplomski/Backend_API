package server

import (
	"diplomski.com/handler"
	"github.com/gin-gonic/gin"
)

func initAuthHandlerRoutes(r *gin.Engine) {
	r.POST("/login", handler.Login)
}

func initProductHandlerRoutes(r *gin.Engine) {
	r.POST("/addProducts", handler.AddProduct)
	r.GET("/getProducts", handler.SearchProducts)
}
