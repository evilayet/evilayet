package county

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
	CountyName string `json:"countyName"`
}

var provinces []Province

func init() {
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
}

func GetCounties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func GetCounty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := mux.Vars(r)["provinceName"]

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(name)) {
			json.NewEncoder(w).Encode(item.ProvinceCounties)
			return
		}
	}

}
