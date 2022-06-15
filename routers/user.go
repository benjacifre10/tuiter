package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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

/* UploadAvatar upload a image from a avatar in the bd */
func UploadAvatarBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	fileType := r.URL.Query().Get("type")
	file, handler, err := r.FormFile(fileType)
	var extension = filepath.Ext(handler.Filename)
	var filePath string = "uploads/" + fileType + "/" + ID + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! " + err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! " + err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	if fileType == "avatar" {
		user.Avatar = ID + extension
	} else {
		user.Banner = ID + extension
	}
	status, err = bd.UpdateUser(user, ID)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el " + fileType + " en la BD ! " + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

/* GetAvatar get the avatar from db and from the directory */
func GetAvatarBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	fileType := r.URL.Query().Get("type")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(fileType) < 1 {
		http.Error(w, "Debe enviar el parametro type", http.StatusBadRequest)
		return
	}

	user, err := bd.FindUser(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	var avatarOrBanner string
	if fileType == "avatar" {
		avatarOrBanner = user.Avatar
	} else {
		avatarOrBanner = user.Banner
	}

	OpenFile, err := os.Open("uploads/" + fileType + "/" + avatarOrBanner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}

/* GetAllUsers give us the list of users */
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.GetAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

