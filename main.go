package main

import (
	"github.com/lyeka/store-test/server"
	"github.com/lyeka/store-test/storage"
	"log"
)

func main() {
	err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}
	server.NewServer().Run()
}
