package api

import (
	"fmt"
	"net/http"
	"time"

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
	accounts := this.router.Group("/accounts")
	{
		accounts.GET("/", this.all)
		accounts.POST("/add", this.add)
	}
}
func (this Account) all(ctx *gin.Context) {
	var rAccounts []models.Account
	var err error
	utils.Query(func(db *storm.DB) {
		err = db.AllByIndex("ID", &rAccounts, storm.Reverse())
		if err != nil {
			fmt.Println("Error al obtener las cuentas")
		} else {
			fmt.Println(&rAccounts)
		}
	})
	ctx.JSON(http.StatusOK, &rAccounts)
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
