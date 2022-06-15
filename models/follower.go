package models

/* struct from a differents user we can folower */
type Follower struct {
	UserId string `bson:"userid" json:"userId"`
	UserFollowerId string `bson:"userfollowerid" json:"userFollowerId"`
}

/* struct response a boolean if the followers exists */
type FollowerResponse struct {
	Status bool `json:"status"`
}
