package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskIDStr := vars["task_id"]

	fmt.Println("task -", taskIDStr)

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

	var msg Response
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusInternalServerError)
		return
	}

	fmt.Println(msg.Title)

	if len(msg.Title) == 0 {
		http.Error(w, "Please fill in Title", http.StatusInternalServerError)
		return
	}

	// response := []byte("POST request received")

	jsonResponse, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)

}

type Person struct {
	Name    string
	Age     int
	Country string
}

type Response struct {
	Task_id int
	Title   string
	Task    string
}
