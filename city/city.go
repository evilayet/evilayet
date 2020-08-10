package city

import (
	"encoding/json"
	"github.com/enesusta/balyoz/text"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

func GetCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")

	log.Println(name)
	log.Println(text.CapitalizeWithTurkish(name))

	for _, item := range cities {
		//if strings.Contains(strings.ToLowerSpecial(unicode.TurkishCase, item.CityName), name) {
		if strings.Contains(item.CityName, text.CapitalizeWithTurkish(name)) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&City{})
}
