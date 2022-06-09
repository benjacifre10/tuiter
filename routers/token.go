package routers

import (
	"errors"
	"strings"

	"github.com/benjacifre10/tuiter/bd"
	"github.com/benjacifre10/tuiter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* Email user will can access everywhere */
var Email string

/* IdUser will can access everywhere */
var IDUser string

/* ProcessToken verify our incoming token with the secret */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	mySecret := []byte("tuiter_benjacifre")
	claims := &models.Claim {}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido") // va sin signos este tipo de errores
	}

	tk = strings.TrimSpace(splitToken[1])

	// recibe el token, lo guarda en claims y el tercer parametro verifica el token con mySecret
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		return claims, false, string(""), err
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	_, find, _ := bd.CheckExistingUser(claims.Email)
	if find == true {
		Email = claims.Email
		IDUser = claims.ID.Hex()
	}

	return claims, find, IDUser, nil
}
