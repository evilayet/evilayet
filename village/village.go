package village

import (
	"encoding/json"
	"github.com/enesusta/balyoz/text"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Province struct {
	ProvinceName     string   `json:"provinceName"`
	ProvinceCounties []County `json:"provinceCounties"`
}

type County struct {
	CountyName  string `json:"countyName"`
	CountyTowns []Town `json:"countyTowns"`
}

type Town struct {
	TownName     string    `json:"townName"`
	TownVillages []Village `json:"townVillages"`
}

type Village struct {
	VillageName string `json:"villageName"`
	ZipCode     int    `json:"zipCode"`
}

var provinces []Province

func init() {
	f, err := os.Open("village.json")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	err = json.Unmarshal(byteValue, &provinces)
}

func GetAllVillages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func GetVillagesOfProvince(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := mux.Vars(r)["provinceName"]

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(provinceName)) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func GetVillagesOfCounty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := mux.Vars(r)["provinceName"]
	countyName := mux.Vars(r)["countyName"]

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(provinceName)) {
			for _, nest := range item.ProvinceCounties {
				if strings.Contains(nest.CountyName, text.CapitalizeWithTurkish(countyName)) {
					json.NewEncoder(w).Encode(nest)
					return
				}
			}
		}
	}
}

func GetVillagesOfTown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := mux.Vars(r)["provinceName"]
	countyName := mux.Vars(r)["countyName"]
	townName := mux.Vars(r)["townName"]

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(provinceName)) {
			for _, nest := range item.ProvinceCounties {
				if strings.Contains(nest.CountyName, text.CapitalizeWithTurkish(countyName)) {
					for _, town := range nest.CountyTowns {
						if strings.Contains(town.TownName, text.CapitalizeWithTurkish(townName)) {
							json.NewEncoder(w).Encode(town)
							return
						}
					}
				}
			}
		}
	}

}
