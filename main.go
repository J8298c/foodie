package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Venue defines struct for venues
type Venue struct {
	Name     string
	Location string
	Phone    string
	Rating   int
	Genre    string
	Dishes   []Dish
}

// Dish defines struct for dishes
type Dish struct {
	Name        string
	Description string
	Price       int
	rating      int
}

func index(w http.ResponseWriter, r *http.Request) {
	// set header
	fmt.Print(" hello index")
	p := "./index.html"
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func getAllVenues(w http.ResponseWriter, r *http.Request) {
	//TODO: share mongo connection dont do it per request
	fmt.Print("fetching all venues")
}

func getSingleVenue(w http.ResponseWriter, r *http.Request) {
	fmt.Print("fetching a venues")
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
