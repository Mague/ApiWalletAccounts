package middlewares

import (
	"github.com/Mague/ApiWalletAccounts/utils"
	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.ValidateToken(ctx)
	}
}
