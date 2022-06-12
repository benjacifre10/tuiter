package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/benjacifre10/tuiter/middlewares"
	"github.com/benjacifre10/tuiter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*
* @func HandlersRoutes
* @desc create the differents routes for my api, and set the port and the permission with cors
 */
func HandlersRoutes() {
	// creo el router a partir de mux
	router := mux.NewRouter()

	router.HandleFunc("/login", middlewares.DbCheck(routers.Login)).Methods("POST")
	router.HandleFunc("/user", middlewares.DbCheck(routers.Register)).Methods("POST")
	router.HandleFunc("/user", middlewares.DbCheck(middlewares.ValidatedJWT(routers.GetUser))).Methods("GET")
	router.HandleFunc("/user", middlewares.DbCheck(middlewares.ValidatedJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.DbCheck(middlewares.ValidatedJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlewares.DbCheck(middlewares.ValidatedJWT(routers.GetTweets))).Methods("GET")

	// traigo el puerto del env, pero si no existe le meto de pecho 8080
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	log.Println("Listening port: " + PORT)
	// creo un handler para que mi api sea accesible desde cualquier lugar
	// con esto tengo los permisos remotos para acceder a esta api, por ejemplo cuando este en heroku
	// le doy permiso a todo el mundo para manejar las rutas que le paso por parametro
	// cors actua como un middleware, porque se fija que permisos tengo y si no los tengo no accedo a las rutas
	handler := cors.AllowAll().Handler(router)

	// con esto pongo mi servidor a escuchar si hubo algun error
	// esto se traduce como http:8080//rutas
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}
