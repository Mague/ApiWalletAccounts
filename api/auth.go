package api

import (
	"fmt"
	"net/http"

	"github.com/Mague/ApiWalletAccounts/models"
	"github.com/Mague/ApiWalletAccounts/utils"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	ctx    *gin.Context
	router *gin.Engine
}

func (this Auth) Load(engine *gin.Engine) {
	this.router = engine
	// this.db = db
	//fmt.Println("Auth success")
	auth := this.router.Group("/auth")
	{
		auth.POST("/sign-in", this.signin)
		auth.GET("/sign-out", this.signout)
	}
}

func (this Auth) signin(ctx *gin.Context) {
	reqUser, reqPwd := ctx.PostForm("userName"), ctx.PostForm("pwd")
	fmt.Println(reqUser, reqPwd)
	uL, pL := len(reqUser), len(reqPwd)
	if (uL > 4 && uL < 15) && pL > 5 && pL < 20 {

		var user models.User
		var err error
		utils.Query(func(db *storm.DB) {
			query := db.Select(q.Eq("UserName", reqUser), q.Eq("Password", reqPwd))
			err = query.First(&user)
			fmt.Println(user)
		})
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Usuario o Contraseña incorrectos",
			})
			return
		}
		user.Password = ""
		token := utils.NewJWT(user)
		result := gin.H{
			"token": token,
			"user": gin.H{
				"Id":       user.ID,
				"UserName": user.UserName,
				"Email":    user.Email,
			},
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func (this Auth) signout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signout Ready!",
	})
}
