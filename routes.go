package main

import (
	"github.com/Mague/ApiWalletAccounts/api"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api.Account{}.Load(router)
	api.User{}.Load(router)
	api.Auth{}.Load(router)
}
