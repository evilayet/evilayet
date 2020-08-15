package town

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetTowns(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/towns", nil)

	GetTowns(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTown(t *testing.T) {

	var res County

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/towns/{provinceName}", nil)

	vars := map[string]string{
		"provinceName": "edi",
	}

	r = mux.SetURLVars(r, vars)

	GetTown(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, mockGetTown())

}

func TestGetSpecificTown(t *testing.T) {
	var res County
	var towns []Town

	towns = append(towns, Town{TownName: "Hasköy"})
	towns = append(towns, Town{TownName: "Havsa"})
	towns = append(towns, Town{TownName: "Merkezköyler"})

	county := County{CountyName: "Havsa", CountyTowns: towns}

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/towns/{provinceName}", nil)

	vars := map[string]string{
		"provinceName": "ed",
		"countyName":   "hav",
	}

	r = mux.SetURLVars(r, vars)

	GetSpecificTown(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, county)

}

func mockGetTown() County {
	var county County
	f, err := os.Open("mock.json")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	err = json.Unmarshal(byteValue, &county)
	return county
}
