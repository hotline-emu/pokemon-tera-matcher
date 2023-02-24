package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client instance
var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(getMongoURI()))
	checkForSimpleError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	checkForSimpleError(err)

	//ping the database
	err = client.Ping(ctx, nil)
	checkForSimpleError(err)

	fmt.Println("Connected to MongoDB")
	defer cancel()

	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("PokemonTeraMatcher").Collection(collectionName)
	return collection
}

func getMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}

func checkForSimpleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
