package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb+srv://root:root@clustersocialmedia.o9seo.mongodb.net/tuitor"

// Exported Mongo Connection
var MongoConnection = ConnectDB() // voy a exportar esta variable que contiene la funcion con el cliente

var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@clustersocialmedia.o9seo.mongodb.net/tuitor?retryWrites=true&w=majority")

/*
* @func ConnectDB
* @desc Release a connection to the DB by a connection string
*
* @return mongoDB client
*/
func ConnectDB() *mongo.Client { // me devuelve una cliente de mongo en un puntero
	// el context es un espacio en memoria que me sirve para setear variables
	// timeouts y es para evitar que si la db se cuelgue, cuelgue la app tambien
	client, err := mongo.Connect(context.TODO(), clientOptions) 

	if err != nil {
		log.Fatal(err.Error())
		return client // lo devuelvo aunque este vacio
	}
	// esto es un ping a la db para ver si esta activa
	err = client.Ping(context.TODO(), nil)
	
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la DB")
	return client
}

/*
* @func CheckConnection
* @desc Ping to the DB to verify its alive
*
* @return int 0(no connection live) | 1(connection live)
*/
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}


