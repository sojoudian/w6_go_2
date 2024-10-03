package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// GET handler
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to the Go API, from")
}

// POST handler
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var msg Message
	// Read body of the POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	// Parse the JSON data
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	// Response back with thesame message
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	fmt.Println("Server running on port: 4455")

	log.Fatal(http.ListenAndServe(":4455", nil))

}
