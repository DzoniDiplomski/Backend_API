package handler

import (
	"database/sql"

	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/service"
	"github.com/DzoniDiplomski/Backend_API/utils"
	"github.com/gin-gonic/gin"
)

var authService = &service.AuthService{}

func Login(c *gin.Context) {
	var authInfo model.Account

	if err := c.BindJSON(&authInfo); err != nil {
		c.Status(400)
		return
	}

	account, err := authService.Login(authInfo)

	if err == sql.ErrNoRows {
		c.String(401, "invalid credentials")
		return
	}
	if err != nil {
		c.String(500, "server error")
		return
	}

	tokenString := utils.GenerateJWT(account.Id)
	c.String(200, tokenString)
}
