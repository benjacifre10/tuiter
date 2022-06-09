package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/benjacifre10/tuiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* FindPerfil find a perfil in the db */
func FindPerfil(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()

	db := MongoConnection.Database("tuiter")
	collection := db.Collection("users")

	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
