package main

import (
	"github.com/enesusta/tzone/city"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/cities", city.GetCities).Methods("GET")
	r.HandleFunc("/city", city.GetCity).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
