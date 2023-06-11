package main

import (
	"flag"
	"log"

	"github.com/MLaskun/go-workshop/api"
	"github.com/MLaskun/go-workshop/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	store, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(*listenAddr, store)
	log.Println("server running on port", *listenAddr)

	log.Fatal(server.Run())

}
