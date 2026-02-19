package main

import (
	"fmt"
	"groupie_tracker/handlers"
	"log"
	"net/http"
)

func main() {
	// Serve static files (CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)

	// Search/filter feature (client-server interaction requirement)
	// http.HandleFunc("/search", handlers.SearchHandler)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
