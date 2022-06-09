package routers

import (
	"encoding/json"
	"net/http"

	"github.com/benjacifre10/tuiter/bd"
)

/* GetPerfil get the user perfil */
func GetPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.FindPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el perfil " + err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
