package models

import "go.mongodb.org/mongo-driver/v2/bson"

type BaseUser struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string        `bson:"username" json:"username"`
	Email      string        `bson:"email" json:"email"`
	EthAccount string        `bson:"ethAccount" json:"ethAccount"`
	Nonce      string        `bson:"nonce" json:"nonce"`
}

type Client struct {
	BaseUser      `bson:",inline"`
	RequestedJobs []bson.ObjectID `bson:"requestedJobs" json:"requestedJobs"`
}

type Freelancer struct {
	BaseUser   `bson:",inline"`
	ActiveJobs []bson.ObjectID `bson:"activeJobs" json:"activeJobs"`
}

type JobStatus string

const (
	UNALLOCATED                 JobStatus = "UNALLOCATED"
	AGREED                      JobStatus = "AGREED"
	CLIENT_STAKED               JobStatus = "CLIENT_STAKED"
	FREELANCER_STAKED           JobStatus = "FREELANCER_STAKED"
	ALL_STAKED_AND_PENDING      JobStatus = "ALL_STAKED_AND_PENDING"
	PENDING_CLIENT_CONFIRMATION JobStatus = "PENDING_CLIENT_CONFIRMATION"
	JOB_COMPLETED               JobStatus = "JOB_COMPLETED"
	CANCEL_REQUESTED            JobStatus = "CANCEL_REQUESTED"
	DEAL_BROKEN                 JobStatus = "DEAL_BROKEN"
	RANDOM_DISPUTED             JobStatus = "RANDOM_DISPUTED"
	PAYMENT_DISPUTED            JobStatus = "PAYMENT_DISPUTED"
)

type Job struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"id"`
	ContractAddress string        `bson:"contractAddress" json:"contractAddress"`
	ClientID        bson.ObjectID `bson:"clientId" json:"clientId"`
	FreelancerID    bson.ObjectID `bson:"freelancerId" json:"freelancerId"`
	Status          JobStatus     `bson:"status" json:"status"`

	// ── Web2 metadata (off-chain, MongoDB-only) ──
	Title        string `bson:"title,omitempty" json:"title"`
	Description  string `bson:"description,omitempty" json:"description"`
	Deadline     string `bson:"deadline,omitempty" json:"deadline"`
	ContactEmail string `bson:"contactEmail,omitempty" json:"contactEmail"`
	PayRange     string `bson:"payRange,omitempty" json:"payRange"`
}
