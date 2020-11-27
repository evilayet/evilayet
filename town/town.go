package town

import (
	"encoding/json"
	"github.com/enesusta/balyoz/text"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type TownEntity struct {
	ProvinceName string
	CountyName   string
}

var provinces []Province
var countiesMap = make(map[string]Province)
var townsMap = make(map[TownEntity]County)

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
	initializeKeyValueMapping(provinces)
}

func initializeKeyValueMapping(provinces []Province) {
	for _, province := range provinces {
		countiesMap[province.ProvinceName] = province
	}

	for key, value := range countiesMap {
		for _, county := range value.ProvinceCounties {
			townEntity := TownEntity{ProvinceName: key, CountyName: county.CountyName}
			townsMap[townEntity] = county
		}
	}
}

func GetTowns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(provinces)
}

func GetTown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := mux.Vars(r)["provinceName"]

	counties, ok := countiesMap[text.CapitalizeWithTurkish(provinceName)]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		_ = json.NewEncoder(w).Encode(counties)
		return
	}

}

func GetSpecificTown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := text.CapitalizeWithTurkish(mux.Vars(r)["provinceName"])
	countyName := text.CapitalizeWithTurkish(mux.Vars(r)["countyName"])

	townEntity := TownEntity{ProvinceName: provinceName, CountyName: countyName}
	town, ok := townsMap[townEntity]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		_ = json.NewEncoder(w).Encode(town)
		return
	}
}
