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

func getAllRestaurants() {
	fmt.Print("Getting all restaurants")
}

func getARestaurant() {
	fmt.Print("Getting a single Restaurant")
}

func main() {
	fmt.Print("hello")
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", r))
}
