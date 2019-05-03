package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./readme.html"))
}

func main() {
	fmt.Print("hello")
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", r))
}
