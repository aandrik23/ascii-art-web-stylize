package main

import (
	"log"
	"net/http"

	"stylize/funcs"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS, JavaScript, etc.)
	staticDir := http.Dir("css")
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(staticDir)))

	staticJSDir := http.Dir("JavaScript")
	mux.Handle("/JavaScript/", http.StripPrefix("/JavaScript/", http.FileServer(staticJSDir)))

	// Handle main page
	mux.HandleFunc("/", funcs.MainPageHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
