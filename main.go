package main

import (
	"github.com/enesusta/tzone/county"
	"github.com/enesusta/tzone/home"
	"github.com/enesusta/tzone/province"
	"github.com/enesusta/tzone/town"
	"github.com/enesusta/tzone/village"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.GetHome).Methods("GET")
	r.HandleFunc("/provinces", province.GetProvinces).Methods("GET")
	r.HandleFunc("/provinces/{provinceName}", province.GetProvince).Methods("GET")
	r.HandleFunc("/counties", county.GetCounties).Methods("GET")
	r.HandleFunc("/counties/{provinceName}", county.GetCounty).Methods("GET")
	r.HandleFunc("/towns", town.GetTowns).Methods("GET")
	r.HandleFunc("/towns/{provinceName}", town.GetTown).Methods("GET")
	r.HandleFunc("/towns/{provinceName}/{countyName}", town.GetSpecificTown).Methods("GET")
	r.HandleFunc("/villages", village.GetAllVillages).Methods("GET")
	r.HandleFunc("/villages/{provinceName}", village.GetVillagesOfProvince).Methods("GET")
	r.HandleFunc("/villages/{provinceName}/{countyName}", village.GetVillagesOfCounty).Methods("GET")
	r.HandleFunc("/villages/{provinceName}/{countyName}/{townName}", village.GetVillagesOfTown).Methods("GET")


	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	log.Fatal(http.ListenAndServe(":12071", c.Handler(r)))
}
