package bd

import (
	"context"
	"log"
	"time"

	"github.com/benjacifre10/tuiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Insert the tweet in the db */
func InsertTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("tweets")

	row := bson.M {
		"userid": t.UserId, 
		"message": t.Message,
		"date": t.Date,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), true, nil
}

/* Get the tweets from the specific user */
func GetTweets(ID string, page int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("tweets")

	var results []*models.Tweet

	condition := bson.M {
		"userid": ID,
	}

	optionsQuery := options.Find()
	optionsQuery.SetLimit(20)
	optionsQuery.SetSort(bson.D {{ Key: "date", Value: -1 }})
	optionsQuery.SetSkip((page -1) * 20)

	tweets, err := collection.Find(ctx, condition, optionsQuery)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for tweets.Next(context.TODO()) {
		var row models.Tweet
		err := tweets.Decode(&row)
		if err != nil {
			return results, false
		}
		results = append(results, &row)
	}

	return results, true
}

/* DeleteTweet delete the tweet from the db */
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M {
		"_id": objID,
		"userid": UserID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}
