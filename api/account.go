package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mague/ApiWalletAccounts/models"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
)

type Account struct {
	ctx    *gin.Context
	router *gin.Engine
	db     *storm.DB
}

func (this Account) Load(engine *gin.Engine) {
	this.router = engine
	// this.db = db
	fmt.Println("rutas cargadas")
	accounts := this.router.Group("/accounts")
	{
		accounts.GET("/", this.all)
		accounts.POST("/add", this.add)
	}
}
func (this Account) all(ctx *gin.Context) {
	db, err := storm.Open("wallet.db")
	// data := models.Account{
	// 	UserName:  "mague",
	// 	Email:     "turronvenezolano@gmail.com",
	// 	Password:  "enmanuel",
	// 	WebSite:   "enmanuelmolina.com",
	// 	CreatedAt: time.Now(),
	// }
	// err = db.Save(&data)
	if err != nil {
		fmt.Println("Error al abrir la base de datos")
	} else {
		fmt.Println("Conexion exitosa")
	}
	var rAccounts []models.Account

	// err = db.Select(q.Eq("UserName", "mague")).Find(&rAccounts)
	err = db.AllByIndex("ID", &rAccounts, storm.Reverse())
	if err != nil {
		fmt.Println("Error al obtener las cuentas")
	} else {
		fmt.Println(&rAccounts)
	}
	db.Close()
	ctx.JSON(http.StatusOK, &rAccounts)
}

func (this Account) add(ctx *gin.Context) {
	fmt.Println("accounts/add")
	db, err := storm.Open("wallet.db")
	if err != nil {
		fmt.Println("Error al abrir la base de datos")
	} else {
		fmt.Println("Conexion exitosa")
	}
	data := models.Account{
		UserName:  ctx.PostForm("userName"),
		Email:     ctx.PostForm("email"),
		Password:  ctx.PostForm("pwd"),
		WebSite:   ctx.PostForm("webSite"),
		CreatedAt: time.Now(),
	}
	err = db.Save(&data)
	db.Close()
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
