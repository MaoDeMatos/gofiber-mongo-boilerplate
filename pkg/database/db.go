package database

import (
	"context"
	"fmt"

	"github.com/maodematos/hi-gofiber/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func GetCollection(name string) *mongo.Collection {
	return MongoClient.Database("test").Collection(name)
}

func StartMongoDB() {
	fmt.Println("🔌 Connecting to database at " + config.Current.MONGODB_URI + "...")

	// var err error
	MongoClient, _ = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(config.Current.MONGODB_URI),
	)

	// Ping the primary
	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("⚠️ Could not connect to primary...")
		panic(err)
	}

	fmt.Println("✨ Successfully connected to database !")
}

func CloseMongoDB() {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("⛔ Successfully closed connection to database !")
}
