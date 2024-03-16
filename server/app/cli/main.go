package main

import (
	"main/api"
	"main/storage"
)

func main() {
	storage := storage.NewDatabase(storage.Config{})

	server := api.New(api.Config{}, *storage)
	panic(server.Start())
}
