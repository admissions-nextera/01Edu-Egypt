package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	StatusCode int
	StatusText string
	Message    string
}

func RenderError(w http.ResponseWriter, statusCode int, message string) {
	// 1. Set the HTTP response status code in the header
	w.WriteHeader(statusCode)

	// 2. Prepare the data for the template
	// http.StatusText(404) returns "Not Found" automatically
	data := ErrorData{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		Message:    message,
	}

	// 3. Parse the error template
	tmpl, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		// If the error template fails to load, fall back to a basic text error
		log.Printf("Crititcal: Error template not found: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 4. Execute the template with ErrorData
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing error template: %v", err)
	}
}
