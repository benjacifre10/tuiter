package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benjacifre10/tuiter/bd"
	"github.com/benjacifre10/tuiter/jwt"
	"github.com/benjacifre10/tuiter/models"
)

/* Login its the door to use the app */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contrasena invalidos " + err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	document, exists := bd.Login(t.Email, t.Password)
	if exists == false {
		http.Error(w, "Usuario y/o contrasena invalidos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token " + err.Error(), 400)
	}

	resp := models.LoginResponse {
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie {
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}
