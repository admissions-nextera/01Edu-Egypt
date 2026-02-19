package api

import (
	"encoding/json"
	"fmt"
	"groupie_tracker/models"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

func FetchData(url string, target any) error {
	// 1. Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}

	// 2. Ensure the body is closed after we are done
	defer resp.Body.Close()

	// 3. Check for successful status codes (200-299 range)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("api returned bad status: %d", resp.StatusCode)
	}

	// 4. Decode directly from the stream (memory efficient)
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}

// GetArtists fetches the full list of artists (returns an array)
func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists, err
}

// GetLocations fetches specific location data for an artist using their unique URL
func GetLocations(url string) (models.Locations, error) {
	var locations models.Locations
	err := FetchData(url, &locations)
	return locations, err
}

// GetDates fetches specific date data for an artist using their unique URL
func GetDates(url string) (models.Dates, error) {
	var dates models.Dates
	err := FetchData(url, &dates)
	return dates, err
}

// GetRelations fetches the map of locations and dates for an artist using their unique URL
func GetRelations(url string) (models.Relation, error) {
	var relation models.Relation
	err := FetchData(url, &relation)
	return relation, err
}
