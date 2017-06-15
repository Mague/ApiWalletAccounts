package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/Mague/ApiWalletAccounts/models"
	"github.com/Mague/ApiWalletAccounts/utils"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			// if token, err := c.Cookie("token"); err == nil || token != "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func signin(ctx *gin.Context) {
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
		token := utils.NewJWT(user, privateKey)

		result := models.Token{Token: token}

		ctx.JSON(http.StatusOK, result)
	}
}

func TestLoginAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)
	LoadSSL()
	r.POST("/auth/sign-in/", ensureNotLoggedIn(), signin)
	loginPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/auth/sign-in", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Test that the page title is "Successful Login"
	// You can carry out a lot more detailed tests using libraries that can
	// parse and process HTML pages
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || p == nil {
		t.Fail()
	}
}
func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("userName", "mague")
	params.Add("pwd", "enmanuel2013")
	return params.Encode()
}

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func LoadSSL() {
	// log.Fatal("init")
	privateBytes, err := ioutil.ReadFile("./ssl/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./ssl/public.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a publicKey")
	}
}
