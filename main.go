package main

import (
	"log"
	"net/http"

	"stylize/funcs" // Adjust this import path as needed
)

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS, JavaScript, images, etc.)
	staticDir := http.Dir("css") // Directory containing static files
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(staticDir)))

	// Register specific routes
	mux.HandleFunc("/", funcs.MainPageHandler)          // Main page handler
	mux.HandleFunc("/ascii-art", funcs.AsciiArtHandler) // ASCII art generation handler

	// Start the server with the custom handler that includes 404 handling
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
