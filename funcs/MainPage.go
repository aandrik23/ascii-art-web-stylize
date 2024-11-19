package funcs

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // Respond with 404 if the path isn't "/"
		return
	}

	// Define the path to the main page template
	templatePath := filepath.Join("Html", "mainpage.html")

	// Parse the template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Println("Error loading template:", err)
		return
	}

	// Execute the template with any data (using `nil` here if there's no data)
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
	}
}
