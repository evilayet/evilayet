package village

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
	TownName     string    `json:"townName"`
	TownVillages []Village `json:"townVillages"`
}

type Village struct {
	VillageName string `json:"villageName"`
	ZipCode     int    `json:"zipCode"`
}

type TownEntity struct {
	ProvinceName string
	CountyName   string
}

type VillageEntity struct {
	ProvinceName string
	CountyName   string
	TownName     string
}

var provinces []Province
var countiesMap = make(map[string][]County)
var townsMap = make(map[TownEntity]County)
var villagesMap = make(map[VillageEntity]Town)

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
	initializeKeyValueMapping(provinces)
}

func initializeKeyValueMapping(provinces []Province) {
	for _, province := range provinces {
		countiesMap[province.ProvinceName] = province.ProvinceCounties
	}

	for key, value := range countiesMap {
		for _, county := range value {
			townEntity := TownEntity{ProvinceName: key, CountyName: county.CountyName}
			townsMap[townEntity] = county
		}
	}

	for key, value := range townsMap {
		for _, town := range value.CountyTowns {
			villageEntity := VillageEntity{ProvinceName: key.ProvinceName, CountyName: key.CountyName, TownName: town.TownName}
			villagesMap[villageEntity] = town
		}
	}
}

func GetAllVillages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func GetVillagesOfProvince(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := text.CapitalizeWithTurkish(mux.Vars(r)["provinceName"])

	counties, ok := countiesMap[provinceName]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		_ = json.NewEncoder(w).Encode(counties)
		return
	}
}

func GetVillagesOfCounty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := text.CapitalizeWithTurkish(mux.Vars(r)["provinceName"])
	countyName := text.CapitalizeWithTurkish(mux.Vars(r)["countyName"])

	townEntity := TownEntity{ProvinceName: provinceName, CountyName: countyName}
	county, ok := townsMap[townEntity]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		_ = json.NewEncoder(w).Encode(county)
		return
	}
}

func GetVillagesOfTown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	provinceName := text.CapitalizeWithTurkish(mux.Vars(r)["provinceName"])
	countyName := text.CapitalizeWithTurkish(mux.Vars(r)["countyName"])
	townName := text.CapitalizeWithTurkish(mux.Vars(r)["townName"])

	villageEntity := VillageEntity{ProvinceName: provinceName, CountyName: countyName, TownName: townName}
	village, ok := villagesMap[villageEntity]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		_ = json.NewEncoder(w).Encode(village)
		return
	}
}
