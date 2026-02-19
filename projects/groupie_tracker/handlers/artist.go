package handlers

import (
	"groupie_tracker/api"
	"groupie_tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ArtistData struct {
	Artist    models.Artist
	Locations models.Locations
	Dates     models.Dates
	Relations models.Relation
}

// ArtistHandler displays individual artist details
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get artist ID from URL query parameter
	idStr := r.URL.Query().Get("id")

	// FIX: Check explicitly for missing id param before Atoi
	if idStr == "" {
		RenderError(w, http.StatusBadRequest, "Missing artist ID")
		return
	}

	targetId, err := strconv.Atoi(idStr)
	if err != nil || targetId < 1 {
		RenderError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	// 2. Fetch all artists
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		RenderError(w, http.StatusInternalServerError, "Failed to fetch artists data")
		return
	}

	// 3. Find the artist with matching ID
	var selectedArtist models.Artist
	found := false
	for _, artist := range artists {
		if targetId == artist.ID {
			selectedArtist = artist
			found = true
			break
		}
	}
	if !found {
		RenderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	// 4. Fetch additional data — FIX: each has its own error check
	locations, err := api.GetLocations(selectedArtist.Locations)
	if err != nil {
		log.Printf("Error fetching locations for artist %d: %v", targetId, err)
		RenderError(w, http.StatusInternalServerError, "Failed to fetch location data")
		return
	}

	dates, err := api.GetDates(selectedArtist.ConcertDates)
	if err != nil {
		log.Printf("Error fetching dates for artist %d: %v", targetId, err)
		RenderError(w, http.StatusInternalServerError, "Failed to fetch date data")
		return
	}

	relations, err := api.GetRelations(selectedArtist.Relations)
	if err != nil {
		log.Printf("Error fetching relations for artist %d: %v", targetId, err)
		RenderError(w, http.StatusInternalServerError, "Failed to fetch relation data")
		return
	}

	// 5. Combine data into a struct for template
	data := ArtistData{
		Artist:    selectedArtist,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}

	// 6. Render artist.html template
	tmpl, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
