package server

import (
	"github.com/DzoniDiplomski/Backend_API/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.Use(middleware.AuthMiddleware)
	initRoutes(r)
	r.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	initAuthHandlerRoutes(r)
	initCashierRoutes(r)
	initManagerRoutes(r)
	initStorageWorkerRoutes(r)
	initDailyMarketRoutes(r)
}
