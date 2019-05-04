package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	// set header
	fmt.Print(" hello index")
	p := "./index.html"
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func getAllVenues(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Getting all venue")
}

func getSingleVenue(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Getting a single venue")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}

func createVenue(w http.ResponseWriter, r *http.Request) {
	fmt.Print("creating a venue")
}

func addDishToVenue(w http.ResponseWriter, r *http.Request) {
	fmt.Print("add a dish to the venue")
}

func editVenue(w http.ResponseWriter, r *http.Request) {
	fmt.Print("edit a venue")
}

func main() {
	fmt.Print("hello")
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/venues/all", getAllVenues).Methods("GET")
	r.HandleFunc("/venues/{id}", getSingleVenue).Methods("GET")
	r.HandleFunc("/venues/new", createVenue).Methods("POST")
	r.HandleFunc("/venues/{venue}/dish", addDishToVenue).Methods("POST")
	r.HandleFunc("/venues/{venue}/edit", editVenue).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
