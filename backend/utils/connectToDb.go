package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DBClient *mongo.Client

func ConnectToDb() (*mongo.Client, error) {
	uri := os.Getenv("MONGO_CONNECT_URI")
	if uri == "" {
		return nil, fmt.Errorf("Empty MongoDB Connection string!")
	}

	DBClient, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("Error Connecting to Database! Error:%s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := DBClient.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Database Non-Responsive! Error:%s", err)
	}

	log.Infof("Database Connection Successful: %s", uri)
	return DBClient, nil
}
