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
