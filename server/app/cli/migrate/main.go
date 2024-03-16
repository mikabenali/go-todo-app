package main

import (
	"context"
	"fmt"
	"main/storage"
	"main/types"
	"os"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Starting migration")

	dbConfig := storage.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Uri:      os.Getenv("MONGO_URI"),
	}
	storage := storage.NewDatabase(dbConfig)
	collection := storage.GetTasksCollection()

	documentsToInsert := []interface{}{
		types.Task{Id: uuid.NewString(), Name: "My first task", Description: "lorem ipsum lalala"},
		types.Task{Id: uuid.NewString(), Name: "My second task", Description: "lorem ipsum lalala"},
		types.Task{Id: uuid.NewString(), Name: "My third task", Description: "lorem ipsum lalala"},
		types.Task{Id: uuid.NewString(), Name: "My other task", Description: "lorem ipsum lalala"},
		types.Task{Id: uuid.NewString(), Name: "My foo task", Description: "lorem ipsum lalala"},
		types.Task{Id: uuid.NewString(), Name: "My bar task", Description: "lorem ipsum lalala"},
	}

	collection.InsertMany(context.TODO(), documentsToInsert)
}
