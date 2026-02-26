package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

const BASEURL = "https://groupietrackers.herokuapp.com/api/"

func fetchArtists() ([]Artist, error) {
	resp, err := http.Get(BASEURL + "artists")
	if err != nil {
		return nil, fmt.Errorf("Can't Get your URL: %v", err)
	}

	defer resp.Body.Close()
	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("Error when Decoding JSON: %v", err)
	}
	return artists, nil
}

func fetchRelations() (Relation, error) {
	resp, err := http.Get(BASEURL + "relation")
	if err != nil {
		return Relation{}, fmt.Errorf("Can't Get your URL: %v", err)
	}

	defer resp.Body.Close()
	var relations Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return Relation{}, fmt.Errorf("Error when Decoding JSON: %v", err)
	}
	return relations, nil
}
