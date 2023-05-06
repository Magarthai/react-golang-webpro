package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	banknotes = append(banknotes, Banknote{ID: "1", Type: "แบบที่1", No: 1, Price: 500})
	banknotes = append(banknotes, Banknote{ID: "2", Type: "แบบที่2", No: 5, Price: 20})
	r.HandleFunc("/api/banknotes", getBanknotes).Methods("GET")
	r.HandleFunc("/api/banknotes/{id}", getBanknote).Methods("GET")
	log.Fatal(http.ListenAndServe((":8000"), r))

}

type Banknote struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	No    int    `json:"no"`
	Price int    `json:"price"`
}

var banknotes []Banknote

// Handle baseurl request
func handle(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(banknotes)
}

// Get all books
func getBanknotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(banknotes)
}

// Get single book
func getBanknote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r) // Get params
	fmt.Println(r)
	// Loop through the books and find with ID
	for _, item := range banknotes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Banknote{})
}

// Create Books
func createBanknote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var banknote Banknote
	_ = json.NewDecoder(r.Body).Decode(&banknote)
	banknote.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe.
	banknotes = append(banknotes, banknote)
	json.NewEncoder(w).Encode(banknote)
}

// Update book
func updateBanknote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range banknotes {
		if item.ID == params["id"] {
			banknotes = append(banknotes[:index], banknotes[index+1:]...)
			var banknote Banknote
			_ = json.NewDecoder(r.Body).Decode(&banknote)
			banknote.ID = params["id"]
			banknotes = append(banknotes, banknote)
			json.NewEncoder(w).Encode(banknote)
			return
		}
	}

	json.NewEncoder(w).Encode(banknotes)
}

// Delete book
func deleteBanknote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range banknotes {
		if item.ID == params["id"] {
			banknotes = append(banknotes[:index], banknotes[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(banknotes)
}
