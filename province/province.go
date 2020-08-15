package province

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
	ProvinceName     string `json:"provinceName"`
}


var provinces []Province

func init() {
	f, err := os.Open("province.json")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("io error %v", err)
	}

	err = json.Unmarshal(byteValue, &provinces)
}

func GetCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func GetCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//name := r.URL.Query().Get("name")
	name := mux.Vars(r)["provinceName"]

	log.Println(name)
	log.Println(text.CapitalizeWithTurkish(name))

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(name)) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Province{})
}
