package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Database

func ConnectMongo() {
	dsnString := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsnString))
	if err != nil {
		log.Fatal("MongoDB Not Connected!")
	}
	dbMongo := client.Database("flox_social")
	MongoClient = dbMongo
}
