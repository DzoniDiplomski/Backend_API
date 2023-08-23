package middleware

import (
	"net/http"
	"strings"

	"github.com/DzoniDiplomski/Backend_API/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	if path := c.Request.URL.Path; path == "/login" {
		c.Next()
		return
	}

	jwtString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	isJWTValid := utils.CheckJWTValidity(jwtString)

	if !isJWTValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func CashierMiddleware(c *gin.Context) {
	role := c.GetHeader("role")

	if role != "KASIR" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do this!"})
		c.Abort()
		return
	}

	c.Next()
}

func ManagerMiddleware(c *gin.Context) {
	role := c.GetHeader("role")

	if role != "POSLOVODJA" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do this!"})
		c.Abort()
		return
	}

	c.Next()
}

func StorageWorkerMiddleware(c *gin.Context) {
	role := c.GetHeader("role")

	if role != "MAGACIONER" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do this!"})
		c.Abort()
		return
	}

	c.Next()
}
