package bd

import (
	"context"
	"time"

	"github.com/benjacifre10/tuiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* Insert one user follower */
func InsertFollower(t models.Follower) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("followers")

	_, err := collection.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}

/* Delete one user follower */
func DeleteFollower(t models.Follower) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("followers")

	_, err := collection.DeleteOne(ctx, t)
	
	if err != nil {
		return false, err
	}

	return true, nil
}

/* Get the followers by userId */
func GetFollowers(t models.Follower) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("followers")

	condition := bson.M {
		"userid": t.UserId,
		"userfollowerid": t.UserFollowerId,
	}

	var result models.Follower

	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}

/* GetTweetsFollowers get the tweets of my followers */
func GetTweetsFollowers(ID string, page int) ([]models.TweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("tuitor")
	collection := db.Collection("followers")

	skip := (page - 1) * 20

	condition := make([]bson.M, 0)
	condition = append(condition, bson.M { "$match": bson.M { "userid": ID } })
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "tweets",
			"localField": "userfollowerid",
			"foreignField": "userid",
			"as": "tweets",
		},
	})
	condition = append(condition, bson.M { "$unwind": "$tweets" })
	condition = append(condition, bson.M { "$sort": bson.M { "tweets.date": -1 }}) // para ordenar va con 1 o -1(asc y desc)
	condition = append(condition, bson.M { "$skip": skip })
	condition = append(condition, bson.M { "$limit": 20 })

	cur, err := collection.Aggregate(ctx, condition)
	var result []models.TweetsFollowers

	err = cur.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
