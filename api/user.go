package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mague/ApiWalletAccounts/models"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
)

type User struct {
	ctx    *gin.Context
	router *gin.Engine
}

func (this User) Load(engine *gin.Engine) {
	this.router = engine
	// this.db = db
	accounts := this.router.Group("/users")
	{
		accounts.GET("/", this.all)
		accounts.POST("/add", this.add)
	}
}

func (this User) all(ctx *gin.Context) {
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
	var rUsers []models.User

	// err = db.Select(q.Eq("UserName", "mague")).Find(&rUsers)
	err = db.AllByIndex("ID", &rUsers, storm.Reverse())
	if err != nil {
		fmt.Println("Error al obtener las cuentas")
	} else {
		fmt.Println(&rUsers)
	}
	db.Close()
	ctx.JSON(http.StatusOK, &rUsers)
}

func (this User) add(ctx *gin.Context) {
	fmt.Println("accounts/add")
	db, err := storm.Open("wallet.db")
	if err != nil {
		fmt.Println("Error al abrir la base de datos")
	} else {
		fmt.Println("Conexion exitosa")
	}
	data := models.User{
		UserName:  ctx.PostForm("userName"),
		Email:     ctx.PostForm("email"),
		Password:  ctx.PostForm("pwd"),
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
