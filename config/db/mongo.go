package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.uber.org/zap"
)

func SetupMongoClient() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	ctx := context.TODO()

	opts := mongoOptions.Client().ApplyURI(uri)
	opts.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	opts.SetReadConcern(readconcern.Majority())
	opts.SetReadPreference(readpref.Primary())
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		zap.S().Fatalf("failure to connect to MongoDB: %s", err)
	}

	zap.S().Debug("mongo db connection setup has been done")
	return client
}
