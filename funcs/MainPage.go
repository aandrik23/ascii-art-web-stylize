package funcs

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// MainPageHandler serves the main page and processes form submissions.
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to 404 if the path is not "/"
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	// Path to the template file
	templatePath := filepath.Join("Html", "mainpage.html")

	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Println("Error loading template:", err)
		return
	}

	// Prepare data for the template
	var data struct {
		Text     string
		Banner   string
		AsciiArt string
		HasError bool
		Error    string
	}

	// Handle POST request (form submission)
	if r.Method == http.MethodPost {
		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form data", http.StatusBadRequest)
			return
		}

		// Retrieve form values
		data.Text = r.FormValue("text")
		data.Banner = r.FormValue("banner")

		// Validate input
		if data.Text == "" || data.Banner == "" {
			data.HasError = true
			data.Error = "Text or banner must be provided."
			w.WriteHeader(http.StatusBadRequest) // Set 400 status
			renderTemplate(w, tmpl, data)
			return
		}

		// Validate ASCII characters in the text
		if !isValidASCII(data.Text) {
			data.HasError = true
			data.Error = "Text contains invalid characters. Only printable ASCII characters are allowed."
			w.WriteHeader(http.StatusBadRequest) // Set 400 status
			renderTemplate(w, tmpl, data)
			return
		}

		// Check if banner file exists
		bannerFilePath := "./banners/" + data.Banner + ".txt"
		file, err := os.ReadFile(bannerFilePath)
		if err != nil {
			data.HasError = true
			if os.IsNotExist(err) {
				data.Error = "Banner file not found."
				w.WriteHeader(http.StatusInternalServerError) // Set 404 status
			} else {
				data.Error = "Internal server error while reading banner file."
				w.WriteHeader(http.StatusInternalServerError) // Set 500 status
			}
			renderTemplate(w, tmpl, data)
			return
		}

		// Process ASCII art generation
		fileContent := strings.ReplaceAll(string(file), "\r\n", "\n")
		lines := strings.Split(fileContent, "\n")
		requestLines := strings.Split(data.Text, "\n")

		asciiArt, err := PrintAsciiArt(requestLines, lines)
		if err != nil {
			data.HasError = true
			data.Error = "Error generating ASCII art: " + err.Error()
			w.WriteHeader(http.StatusInternalServerError) // Set 500 status
		} else {
			data.AsciiArt = asciiArt
		}
	}

	// Render the template with data
	renderTemplate(w, tmpl, data)
}

// Helper function to render templates
func renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
	}
}
