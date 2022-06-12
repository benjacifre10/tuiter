package bd

import (
	"context"
	"log"
	"time"

	"github.com/benjacifre10/tuiter/models"
	"github.com/benjacifre10/tuiter/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Register user in bd */
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("users")

	u.Password, _ = utils.EncryptPassword(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.Hex(), true, nil
}

/* Check if an user already existing in db */
func CheckExistingUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("users")

	condition := bson.M{"email" : email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

/* Login the user to the app */
func Login(email string, password string) (models.User, bool) {
	user, find, _ := CheckExistingUser(email)
	if find == false {
		return user, false
	}

	err := utils.DecryptPassword(user.Password, password)
	if err != nil {
		return user, false
	}
	return user, true
}

/* FindUser find a user in the db */
func FindUser(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("users")

	var user models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&user)
	user.Password = ""

	if err != nil {
		log.Println("Registro no encontrado " + err.Error())
		return user, err
	}

	return user, nil
}

/* UpdateUser update the user in the db */
func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("users")

	row := make(map[string]interface{})
	if len(u.Name) > 0 {
		row["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		row["surname"] = u.Surname
	}
	row["birthday"] = u.Birthday
	if len(u.Banner) > 0 {
		row["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		row["biography"] = u.Biography
	}
	if len(u.Ubication) > 0 {
		row["ubication"] = u.Ubication
	}
	if len(u.Website) > 0 {
		row["website"] = u.Website
	}
	if len(u.Avatar) > 0 {
		row["avatar"] = u.Avatar
	}

	updateString := bson.M {
		"$set": row,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
