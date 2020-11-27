package county

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
	CountyName string `json:"countyName"`
}

var countiesMap = make(map[string][]County)

func init() {
	var provinces []Province
	f, err := os.Open("county.json")

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
	for _, item := range provinces {
		countiesMap[item.ProvinceName] = item.ProvinceCounties
	}
}

func GetCounties(w http.ResponseWriter, r *http.Request) {
	var provinces []Province
	w.Header().Set("Content-Type", "application/json")

	for key, value := range countiesMap {
		provinces = append(provinces, Province{ProvinceName: key, ProvinceCounties: value})
	}

	json.NewEncoder(w).Encode(provinces)
}

func GetCounty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := mux.Vars(r)["provinceName"]

	counties, ok := countiesMap[text.CapitalizeWithTurkish(name)]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		_ = json.NewEncoder(w).Encode(counties)
		return
	}
}
