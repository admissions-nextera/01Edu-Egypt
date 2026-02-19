package api

import (
	"testing"
)

func TestGetArtists(t *testing.T) {
	artists, err := GetArtists()

	if err != nil {
		t.Fatalf("Expected no error fetching artists, got: %v", err)
	}

	if len(artists) == 0 {
		t.Fatal("Expected at least one artist, got empty slice")
	}

	// Validate first artist has required fields
	first := artists[0]

	if first.ID == 0 {
		t.Error("Expected artist to have a non-zero ID")
	}
	if first.Name == "" {
		t.Error("Expected artist to have a name")
	}
	if first.Image == "" {
		t.Error("Expected artist to have an image URL")
	}
	if len(first.Members) == 0 {
		t.Error("Expected artist to have at least one member")
	}
	if first.Locations == "" {
		t.Error("Expected artist to have a locations URL")
	}
	if first.ConcertDates == "" {
		t.Error("Expected artist to have a concertDates URL")
	}
	if first.Relations == "" {
		t.Error("Expected artist to have a relations URL")
	}
}

func TestGetLocations(t *testing.T) {
	// First get an artist to get a real locations URL
	artists, err := GetArtists()
	if err != nil || len(artists) == 0 {
		t.Fatal("Could not fetch artists to test locations")
	}

	locations, err := GetLocations(artists[0].Locations)
	if err != nil {
		t.Fatalf("Expected no error fetching locations, got: %v", err)
	}

	if locations.ID == 0 {
		t.Error("Expected locations to have a non-zero ID")
	}
	if len(locations.Locations) == 0 {
		t.Error("Expected at least one location")
	}
}

func TestGetDates(t *testing.T) {
	artists, err := GetArtists()
	if err != nil || len(artists) == 0 {
		t.Fatal("Could not fetch artists to test dates")
	}

	dates, err := GetDates(artists[0].ConcertDates)
	if err != nil {
		t.Fatalf("Expected no error fetching dates, got: %v", err)
	}

	if dates.ID == 0 {
		t.Error("Expected dates to have a non-zero ID")
	}
	if len(dates.Dates) == 0 {
		t.Error("Expected at least one date")
	}
}

func TestGetRelations(t *testing.T) {
	artists, err := GetArtists()
	if err != nil || len(artists) == 0 {
		t.Fatal("Could not fetch artists to test relations")
	}

	relations, err := GetRelations(artists[0].Relations)
	if err != nil {
		t.Fatalf("Expected no error fetching relations, got: %v", err)
	}

	if relations.ID == 0 {
		t.Error("Expected relations to have a non-zero ID")
	}
	if len(relations.DatesLocations) == 0 {
		t.Error("Expected at least one dates-locations entry")
	}
}

func TestFetchDataBadURL(t *testing.T) {
	var target interface{}
	err := FetchData("https://groupietrackers.herokuapp.com/api/nonexistent", &target)

	if err == nil {
		t.Error("Expected an error for a bad URL, got nil")
	}
}
