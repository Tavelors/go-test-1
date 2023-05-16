package pack

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SayHello() {
	fmt.Println("Hello from the utility function!")
}

type Response struct {
	Message string `json:"message"`
}

func hi(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}
