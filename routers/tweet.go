package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/benjacifre10/tuiter/bd"
	"github.com/benjacifre10/tuiter/models"
)

/* InsertTweet let us save the tweet in the db */
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	row := models.Tweet {
		UserId: IDUser,
		Message: tweet.Message,
		Date: time.Now(),
	}

	_, status, err := bd.InsertTweet(row)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el tweet, reintente nuevamente " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* GetTweets let get all the tweets from a specific user */
func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro de pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro de pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	response, correct := bd.GetTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
