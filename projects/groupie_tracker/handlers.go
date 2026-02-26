package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ArtistData struct {
	Artist   Artist
	Relation map[string][]string
}

var artists []Artist
var relations Relation

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 400)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", 404)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if err := tmpl.Execute(w, artists); err != nil {
		log.Printf("template execute error: %v", err)
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 400)
		return
	}

	id := r.URL.Query().Get("id")
	idN, err := strconv.Atoi(id)
	if err != nil || idN < 1 {
		http.Error(w, "Missing or invalid ID", 400)
		return
	}

	var found *Artist
	for i, a := range artists {
		if a.ID == idN {
			found = &artists[i]
			break
		}
	}
	fmt.Println(found)
	if found == nil {
		http.NotFound(w, r)
		return
	}

	var foundRelation map[string][]string
	for _, rel := range relations.Index {
		if rel.ID == idN {
			foundRelation = rel.DatesLocations
			break
		}
	}

	data := ArtistData{
		Artist:   *found,
		Relation: foundRelation,
	}

	tmpl, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("template execute error: %v", err)
	}
}
