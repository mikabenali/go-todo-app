package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

// Not real migrations, just a data set to test the api
func main() {
	fmt.Println("Starting migration")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	defer func() {
		client.Disconnect(context.TODO())
	}()

	collection := client.Database("todo").Collection("tasks")
	collection.DeleteMany(context.TODO(), bson.M{})

	documentsToInsert := []interface{}{
		Task{Name: "My first task", Description: "lorem ipsum"},
		Task{Name: "My second task", Description: "lorem ipsum"},
		Task{Name: "My third task", Description: "lorem ipsum"},
		Task{Name: "My other task", Description: "lorem ipsum"},
		Task{Name: "My foo task", Description: "lorem ipsum"},
		Task{Name: "My bar task", Description: "lorem ipsum"},
	}

	if _, err := collection.InsertMany(context.TODO(), documentsToInsert); err != nil {
		panic(err)
	}
}
