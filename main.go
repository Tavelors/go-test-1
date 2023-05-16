package main

import (
	"encoding/json"
	"go-test-1/pack"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello, World!",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
\\
		w.Write(jsonResponse)
	}

	pack.SayHello()

	http.HandleFunc("/api", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
