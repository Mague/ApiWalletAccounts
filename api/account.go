package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mague/ApiWalletAccounts/middlewares"
	"github.com/Mague/ApiWalletAccounts/models"
	"github.com/Mague/ApiWalletAccounts/utils"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
)

type Account struct {
	ctx    *gin.Context
	router *gin.Engine
}

func (this Account) Load(engine *gin.Engine) {
	this.router = engine
	// this.db = db
	accounts := this.router.Group("/accounts", middlewares.EnsureLoggedIn())
	{
		accounts.GET("/:id", this.get)
		accounts.POST("/add", this.add)
	}
}
func (this Account) get(ctx *gin.Context) {
	var rAccounts []models.Account
	var err error
	id := ctx.Params.ByName("id")
	utils.Query(func(db *storm.DB) {
		err = db.Find("CurrentUser", id, &rAccounts, storm.Reverse())
		if err != nil {
			fmt.Println("Error al obtener las cuentas")
		} else {
			fmt.Println(&rAccounts)
		}
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	} else {
		ctx.JSON(http.StatusOK, &rAccounts)
	}
}

func (this Account) add(ctx *gin.Context) {
	var err error
	data := models.Account{
		UserName:  ctx.PostForm("userName"),
		Email:     ctx.PostForm("email"),
		Password:  ctx.PostForm("pwd"),
		WebSite:   ctx.PostForm("webSite"),
		CreatedAt: time.Now(),
	}
	utils.Query(func(db *storm.DB) {
		err = db.Save(&data)
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Error al añadir al usuario",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Usuario añadido correctamente",
		})
	}
}
