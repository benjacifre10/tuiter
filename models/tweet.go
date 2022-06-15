package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Tweet model for the mongo DB */
type Tweet struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId string `bson:"userid" json:"userId,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date time.Time `bson:"date" json:"date,omitempty"`
}

/* Return the tweets from the differents followers */
type TweetsFollowers struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId string `bson:"userid" json:"userId,omitempty"`
	UserFollowerId string `bson:"userfollowerid" json:"userFollowerId,omitempty"`
	Tweets struct {
		Message string `bson:"message" json:"message,omitempty"`
		Date time.Time `bson:"date" json:"date,omitempty"`
		ID string `bson:"_id" json:"_id,omitempty"`
	}
}
