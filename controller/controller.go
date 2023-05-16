package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Person{
		Name:    "shaun",
		Age:     27,
		Country: "uk",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func PostTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fmt.Println(body)

	response := []byte("POST request received")
	w.Write(response)
}

type Person struct {
	Name    string
	Age     int
	Country string
}
