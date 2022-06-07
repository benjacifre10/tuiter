package main

import (
	"log"

	"github.com/benjacifre10/tuiter/bd"
	"github.com/benjacifre10/tuiter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}

	handlers.HandlersRoutes()
}
