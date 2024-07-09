package config

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client  {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(LoadFromURIFromENV()))
    if err != nil {
        log.Fatal(err)
    }

	// Check the connection

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
    fmt.Println("Connected to MongoDB")
    return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("golangAPI").Collection(collectionName)
    return collection
}