package main

import (
	"main/api"
	"main/storage"
	"os"
)

func main() {
	dbConfig := storage.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Uri:      os.Getenv("MONGO_URI"),
	}
	storage := storage.NewDatabase(dbConfig)

	server := api.New(api.Config{}, *storage)
	panic(server.Start())
}
