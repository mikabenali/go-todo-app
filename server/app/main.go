package main

import (
	"main/api"
	"main/storage"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("MONGO_URI") == "" {
		godotenv.Load()
	}

	dbConfig := storage.Config{
		Uri: os.Getenv("MONGO_URI"),
	}
	storage := storage.NewDatabase(dbConfig)

	server := api.New(api.Config{}, *storage)
	panic(server.Start())
}
