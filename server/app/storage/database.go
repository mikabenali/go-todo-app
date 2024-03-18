package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Uri string
}

type Database struct {
	Config Config
	Client *mongo.Client
}

func NewDatabase(config Config) *Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Uri))
	if err != nil {
		panic(err)
	}

	return &Database{
		Config: config,
		Client: client,
	}
}

func (db *Database) GetDatabase() *mongo.Database {
	return db.Client.Database("todo")
}
