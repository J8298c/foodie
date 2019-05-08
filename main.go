package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost/yumgum"

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
	fmt.Print("Getting all venue")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("yumgum").Collection("venues")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "results: %v\n", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
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
