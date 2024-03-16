package storage

import (
	"context"
	"main/types"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Username string
	Password string
	Uri      string
}

type Database struct {
	Config Config
	Client *mongo.Client

	// For testing
	Tasks []types.Task
}

func NewDatabase(config Config) *Database {
	tasks := []types.Task{
		{Id: uuid.NewString(), Name: "My first task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My second task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My third task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My other task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My foo task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My bar task", Description: "lorem ipsum lalala"},
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Uri))
	if err != nil {
		panic(err)
	}

	return &Database{
		Config: config,
		Client: client,

		Tasks: tasks,
	}
}

func (db *Database) GetDatabase() *mongo.Database {
	return db.Client.Database("todo")
}
