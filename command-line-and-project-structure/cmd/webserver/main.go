package main

import (
	poker "command-line-and-project-structure"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemFileStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("could not listen on port 4000, %v", err)
	}
}
