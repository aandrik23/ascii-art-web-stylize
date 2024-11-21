package funcs

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
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

		// Check for missing values
		if data.Text == "" || data.Banner == "" {
			data.HasError = true
			data.Error = "Both text and banner must be provided."
		} else {
			// Check if banner file exists
			bannerFilePath := "./banners/" + data.Banner + ".txt"
			file, err := os.ReadFile(bannerFilePath)
			if err != nil {
				data.HasError = true
				if os.IsNotExist(err) {
					data.Error = "Banner file not found."
				} else {
					data.Error = "Internal server error while reading banner file."
				}
			} else {
				// Process the ASCII art generation
				fileContent := strings.ReplaceAll(string(file), "\r\n", "\n")
				lines := strings.Split(fileContent, "\n")
				requestLines := strings.Split(data.Text, "\\n")

				asciiArt, err := PrintAsciiArt(requestLines, lines)
				if err != nil {
					data.HasError = true
					data.Error = "Error generating ASCII art: " + err.Error()
				} else {
					data.AsciiArt = asciiArt
				}
			}
		}
	}

	// Render the template with the data
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
	}
}
