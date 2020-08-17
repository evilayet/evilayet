package home

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var doc []byte

func init() {
	f, err := os.Open("README.md")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	doc = byteValue
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/markdown; charset=UTF-8")
	_, _ = w.Write(doc)
}
