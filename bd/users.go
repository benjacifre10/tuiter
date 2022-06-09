package bd

import (
	"context"
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
