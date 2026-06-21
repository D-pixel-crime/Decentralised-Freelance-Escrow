package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DBClient *mongo.Client

func createUserIndexes() error {
	coll := DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("client")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}, {Key: "email", Value: 1}, {Key: "ethAccount", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	coll = DBClient.Database(os.Getenv("DATABASE_NAME")).Collection("freelancer")
	indexModel = mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}, {Key: "email", Value: 1}, {Key: "ethAccount", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = coll.Indexes().CreateOne(context.TODO(), indexModel)

	return err
}

func ConnectToDb() (*mongo.Client, error) {
	uri := os.Getenv("MONGO_CONNECT_URI")
	if uri == "" {
		return nil, fmt.Errorf("Empty MongoDB Connection String!")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("Error Connecting to Database! Error:%s", err)
	}

	DBClient = client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := DBClient.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Database Non-Responsive! Error:%s", err)
	}

	err = createUserIndexes()
	if err != nil {
		log.Fatalf("Error creating Database Indexes! Error:%s", err)
		return DBClient, err
	}

	log.Infof("Database Connection Successful: %s", uri)
	return DBClient, nil
}
