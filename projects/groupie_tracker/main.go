package main

import (
	"log"
	"net/http"
)

func init() {
	var err error
	artists, err = fetchArtists()
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	relations, err = fetchRelations()
	if err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
