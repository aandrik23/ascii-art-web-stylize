package funcs

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Check for required fields
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		http.Error(w, "Missing text or banner parameter", http.StatusBadRequest)
		return
	}

	// Check if the banner file exists
	bannerFilePath := "./banners/" + banner + ".txt"
	file, err := os.ReadFile(bannerFilePath)
	if err != nil {
		// If the error is due to file not being found, return 404
		if os.IsNotExist(err) {
			http.Error(w, "Banner file not found", http.StatusNotFound)
		} else {
			// Any other error, like permissions or file system issues, should be 500
			http.Error(w, "Internal server error while reading banner file", http.StatusInternalServerError)
		}
		return
	}

	// Process the banner file content
	fileContent := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(fileContent, "\n")
	requestLines := strings.Split(text, "\n")

	// Generate ASCII art
	asciiArt, err := PrintAsciiArt(requestLines, lines)
	if err != nil {
		http.Error(w, "Error generating ASCII art: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the generated ASCII art as plain text (no <pre> tags)

	fmt.Fprint(w, asciiArt)
}
