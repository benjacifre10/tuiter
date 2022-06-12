package routers

import (
	"encoding/json"
	"net/http"

	"github.com/benjacifre10/tuiter/bd"
	"github.com/benjacifre10/tuiter/models"
)

func Register (w http.ResponseWriter, r *http.Request) {
	var t models.User
	// el r.Body funciona una sola vez, se lee y luego se destruye
	err := json.NewDecoder(r.Body).Decode(&t)
  if err != nil {
		http.Error(w, "Error en los datos recibidos " + err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener al menos 6 caracteres", 400)
		return
	}

	_, findUser, _ := bd.CheckExistingUser(t.Email)

	if findUser == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	userId, status, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al registrar el usuario " + err.Error(), 400)
		return 
	}
	
	// esto es un error interno de mongo que no te inserta nada
	if status == false {
		http.Error(w, "No se ha logrado registrar el usuario", 400)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "El usuario se ha registrado correctamente"
	resp["id"] = userId

	// esto es como devolver un 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

/* GetUser get the user info */
func GetUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	user, err := bd.FindUser(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar los datos del usuario " + err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

/* UpdateUser update the user info */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos " + err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.UpdateUser(t, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el usuario. Reintente nuevamente " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
