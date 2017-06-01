package utils

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/Mague/ApiWalletAccounts/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func NewJWT(user models.User, privateKey *rsa.PrivateKey) string {
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
