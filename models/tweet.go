package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Tweet model for the mongo DB */
type Tweet struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId string `bson:"userid" json:"userid,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date time.Time `bson:"date" json:"date,omitempty"`
}
