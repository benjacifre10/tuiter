package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/benjacifre10/tuiter/bd"
	"github.com/benjacifre10/tuiter/models"
)

/* Add a follower user */
func InsertFollower(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Follower
	t.UserId = IDUser
	t.UserFollowerId = ID

	status, err := bd.InsertFollower(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar agregar un seguidor " + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el seguidor", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* Delete a follower user */
func DeleteFollower(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Follower
	t.UserId = IDUser
	t.UserFollowerId = ID

	status, err := bd.DeleteFollower(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar eliminar un seguidor " + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado eliminar el seguidor", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

/* Get the followers */
func GetFollower(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Follower
	t.UserId = IDUser
	t.UserFollowerId = ID

	var res models.FollowerResponse
	res.Status = true

	status, err := bd.GetFollowers(t)
	if err != nil || status == false {
		res.Status = false
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/* Get the tweets of the followers */
func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "El parametro pagina debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	result, status := bd.GetTweetsFollowers(IDUser, page)
	if status == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
