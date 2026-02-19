package models

// Artist represents a band/artist from the /api/artists endpoint
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

// Locations represents an individual location entry
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Dates represents an individual date entry
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relation links locations and dates via a map
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// FIX: Each index wrapper must be its own struct with its own json tag.
// Previously all three fields shared "index" which caused unmarshal conflicts.

// LocationsIndex holds the wrapper for the locations endpoint
type LocationsIndex struct {
	Index []Locations `json:"index"`
}

// DatesIndex holds the wrapper for the dates endpoint
type DatesIndex struct {
	Index []Dates `json:"index"`
}

// RelationIndex holds the wrapper for the relations endpoint
type RelationIndex struct {
	Index []Relation `json:"index"`
}
