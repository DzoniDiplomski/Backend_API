package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.Use(configCORS())
	// r.Use(middleware.AuthMiddleware)
	initRoutes(r)
	r.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	initAuthHandlerRoutes(r)
	initCashierRoutes(r)
	initManagerRoutes(r)
	initStorageWorkerRoutes(r)
	initDailyMarketRoutes(r)
	initCashierManagerRoutes(r)
}

func configCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	return cors.New(config)
}
