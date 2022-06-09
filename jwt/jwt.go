package jwt

import (
	"time"

	"github.com/benjacifre10/tuiter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* GenerateJWT generated the encrypt token with JWT */
func GenerateJWT(t models.User) (string, error) {
	mySecret := []byte("tuiter_benjacifre")

	payload := jwt.MapClaims {
		"email": t.Email,
		"name": t.Name,
		"surname": t.Surname,
		"birthday": t.Birthday,
		"biography": t.Biography,
		"ubication": t.Ubication,
		"website": t.Website,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),// unix me lo transforma en un formato long
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
