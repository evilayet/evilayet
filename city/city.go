package city

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type City struct {
	CityID   string `json:"cityId"`
	CityName string `json:"cityName"`
}

var cities []City

func init() {
	f, err := os.Open("city.json")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("io error %v", err)
	}

	err = json.Unmarshal(byteValue, &cities)
}

func GetCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}
