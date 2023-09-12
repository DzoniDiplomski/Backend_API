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
	jwtString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	claims, err := utils.DecodeJWT(jwtString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to decode the jwt token! Error: " + err.Error()})
		c.Abort()
		return
	}

	role := claims["role"]

	if role != "KASIR" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do this!"})
		c.Abort()
		return
	}

	c.Next()
}

func CashierManagerMiddleware(c *gin.Context) {
	jwtString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	claims, err := utils.DecodeJWT(jwtString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to decode the jwt token! Error: " + err.Error()})
		c.Abort()
		return
	}

	role := claims["role"]

	if role == "KASIR" || role == "MENADZER" {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do this!"})
	c.Abort()
}

func ManagerMiddleware(c *gin.Context) {
	jwtString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	claims, err := utils.DecodeJWT(jwtString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to decode the jwt token! Error: " + err.Error()})
		c.Abort()
		return
	}

	role := claims["role"]
	if role != "MENADZER" {
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
