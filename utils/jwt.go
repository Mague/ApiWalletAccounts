package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Mague/ApiWalletAccounts/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
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

func NewJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Seguridad!!!",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Error al firmar el token")
	}

	return result
}

func ValidateToken(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				ctx.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("Su token ha expirado")
				return
			case jwt.ValidationErrorSignatureInvalid:
				ctx.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("La firma del token no coincide")
				return
			default:
				ctx.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("Su token no es valido")
				return
			}
		default:
			//ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("Su token no es valido")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return

		}
	}

	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		fmt.Println("Bienvenido al sistema")
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("Su token no es valido")
	}
}
