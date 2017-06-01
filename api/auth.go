package api

import (
	"fmt"
	"net/http"

	"github.com/Mague/ApiWalletAccounts/models"
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
		db, err := storm.Open("wallet.db")
		var user models.User
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Error al establecer conexión con la base de datos",
			})
			return
		}
		// var sUser models.User
		// sUser.UserName = user
		// sUser.UserName = pwd
		query := db.Select(q.Eq("UserName", reqUser), q.Eq("Password", reqPwd))
		err = query.First(&user)
		fmt.Println(user)
		db.Close()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Usuario o Contraseña incorrectos",
			})
			return
		}
		user.Password = ""
		ctx.JSON(http.StatusOK, &user)
	}
}
func (this Auth) signout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signout Ready!",
	})
}
