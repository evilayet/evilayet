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

func GetProvinces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func GetProvince(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//provinceName := r.URL.Query().Get("provinceName")
	provinceName := mux.Vars(r)["provinceName"]

	log.Println(provinceName)
	log.Println(text.CapitalizeWithTurkish(provinceName))

	for _, item := range provinces {
		if strings.Contains(item.ProvinceName, text.CapitalizeWithTurkish(provinceName)) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
