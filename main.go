package main

import (
	"log"
	"net/http"
	"profile_api/internals/handlers"
)

func main() {
	servemux := http.NewServeMux()
	servemux.HandleFunc("/me", handlers.ProfileHandler)
	log.Printf("Starting server on :8080")
	http.ListenAndServe(":8080", servemux)
}
