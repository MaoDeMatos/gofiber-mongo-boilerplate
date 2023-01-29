package database

import (
	"fmt"

	"github.com/maodematos/hi-gofiber/config"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func StartMongoDB() {
	fmt.Println("🔌 Connecting to database at " + config.Current.MONGO_URI + "...")

	if err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI(config.Current.MONGO_URI)); err != nil {
		fmt.Println("⚠️ Could not connect to database...")
		panic(err)
	}

	var currentDb *mongo.Database
	_, MongoClient, currentDb, _ = mgm.DefaultConfigs()

	fmt.Println("✨ Successfully connected to database ! (" + currentDb.Name() + ")")
}

func CloseMongoDB() {
	if err := MongoClient.Disconnect(mgm.Ctx()); err != nil {
		panic(err)
	}
	fmt.Println("⛔ Successfully closed connection to database !")
}
