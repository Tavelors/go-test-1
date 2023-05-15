package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Define the HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello, World!",
		}

		// Convert the response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		w.Write(jsonResponse)
	}

	// Register the handler function to handle requests at the /api endpoint
	http.HandleFunc("/api", handler)

	// Start the HTTP server on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
