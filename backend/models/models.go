package models

import "go.mongodb.org/mongo-driver/v2/bson"

type BaseUser struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string        `bson:"username" json:"username"`
	Email      string        `bson:"email" json:"email"`
	Password   string        `bson:"password" json:"-"`
	EthAccount string        `bson:"account" json:"account"`
}

type Client struct {
	BaseUser      `bson:",inline"`
	RequestedJobs []bson.ObjectID `bson:"requestedJobs" json:"requestedJobs"`
}

type Freelancer struct {
	BaseUser   `bson:",inline"`
	ActiveJobs []bson.ObjectID `bson:"activeJobs" json:"activeJobs"`
}

type Job struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	ClientID     bson.ObjectID `bson:"clientId" json:"clientId"`
	FreelancerID bson.ObjectID `bson:"freelancerId" json:"freelancerId"`
	Status       string        `bson:"status" json:"status"`
}
