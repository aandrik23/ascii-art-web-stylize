package funcs

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// AsciiArtHandler handles ASCII art generation requests.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
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

	// Validate input (only allow ASCII printable characters and newline)
	if !isValidASCII(text) {
		http.Error(w, "Text contains invalid characters. Only printable ASCII characters are allowed.", http.StatusBadRequest)
		return
	}

	// Check if the banner file exists
	bannerFilePath := "./banners/" + banner + ".txt"
	file, err := os.ReadFile(bannerFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Banner file not found", http.StatusNotFound)
		} else {
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

	// Send the generated ASCII art as plain text
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, asciiArt)
}

// Helper function to validate input as printable ASCII (32-126 or newlines)
func isValidASCII(s string) bool {
	for _, c := range s {
		if (c < 32 || c > 126) && c != '\n' {
			return false
		}
	}
	return true
}
