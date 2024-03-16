package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Starting migration")

	connString := fmt.Sprintf(
		"mongodb://%s:%s@mongodb:27017/",
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}

	collection := client.Database("todo").Collection("tasks")

	documentsToInsert := []interface{}{
		Task{Name: "My first task", Description: "lorem ipsum lalala"},
		Task{Name: "My second task", Description: "lorem ipsum lalala"},
		Task{Name: "My third task", Description: "lorem ipsum lalala"},
		Task{Name: "My other task", Description: "lorem ipsum lalala"},
		Task{Name: "My foo task", Description: "lorem ipsum lalala"},
		Task{Name: "My bar task", Description: "lorem ipsum lalala"},
	}

	collection.InsertMany(context.TODO(), documentsToInsert)
}
