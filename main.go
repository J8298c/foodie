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
}

func createVenue() {
	fmt.Print("creating a venue")
}

func addDishToVenue() {
	fmt.Print("add a dish to the venue")
}

func editVenue() {
	fmt.Print("edit a venue")
}

func main() {
	fmt.Print("hello")
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/venues/all", getAllVenues)
	r.HandleFunc("/venues/{id}", getSingleVenue)
	log.Fatal(http.ListenAndServe(":8000", r))
}
