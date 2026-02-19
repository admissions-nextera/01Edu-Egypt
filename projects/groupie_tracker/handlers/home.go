package handlers

import (
	"groupie_tracker/api"
	"html/template"
	"log"
	"net/http"
)

// HomeHandler displays all artists
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Check if path is exactly "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 2. Fetch artists from API
	artists, err := api.GetArtists()
	if err != nil {
		// Log the actual error for the developer, send a generic one to the user
		log.Printf("Error fetching artists: %v", err)
		RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// 3. Parse HTML template
	// Use a relative path that matches your project structure
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	// 4. Execute template with artists data
	// Note: No need for &artists, passing the slice directly is standard
	err = tmpl.Execute(w, artists)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
