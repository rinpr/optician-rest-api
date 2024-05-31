package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectToMongo() *mongo.Client {
	// mongodb connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Getting username and password from .env
	var env map[string]string
	env, _ = godotenv.Read()
	username := env["MONGO_DB_USERNAME"]
	password := env["MONGO_DB_PASSWORD"]

	// Setting username and password
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// Connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check if database is working or not
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB :(")
		log.Fatal(err)
		return nil
	}
	fmt.Println("Successfully connected to MongoDB :)")
	return client
}

var client *mongo.Client = ConnectToMongo()

func GetBillsCollection() *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("optical-bills-data").Collection("data")
	return productCollection
}
