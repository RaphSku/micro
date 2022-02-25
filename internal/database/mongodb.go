package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(mongodb_uri string) (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func CreateDatabase(client *mongo.Client, name string) *mongo.Database {
	return client.Database(name)
}

func CreateCollection(database *mongo.Database, name string) *mongo.Collection {
	return database.Collection(name)
}

func PrintAvailableDatabases(ctx context.Context, client *mongo.Client) {
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
