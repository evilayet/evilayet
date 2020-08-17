package town

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
	TownName string `json:"townName"`
}

var provinces []Province

func init() {
	f, err := os.Open("towns.json")

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

func GetTowns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func GetTown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := mux.Vars(r)["provinceName"]

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(provinceName)) {
			json.NewEncoder(w).Encode(item.ProvinceCounties)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func GetSpecificTown(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusNotFound)
}
