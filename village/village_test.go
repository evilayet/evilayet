package village

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

func TestGetAllVillages(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/provinces", nil)

	GetAllVillages(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetVillagesOfProvince(t *testing.T) {

	var res Province

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/villages/{provinceName}", nil)

	vars := map[string]string{
		"provinceName": "edirne",
	}

	r = mux.SetURLVars(r, vars)

	GetVillagesOfProvince(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, mockGetVillagesOfProvince())
}

func TestGetVillagesOfCounty(t *testing.T) {
	var res County

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/villages/{provinceName}/{countyName}", nil)

	vars := map[string]string{
		"provinceName": "edirne",
		"countyName":   "merkez",
	}

	r = mux.SetURLVars(r, vars)

	GetVillagesOfCounty(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, mockGetVillagesOfCounty())
}

func TestGetVillagesOfTown(t *testing.T) {
	t.Parallel()
	var res Town

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/villages/{provinceName}/{countyName}/{townName}", nil)

	vars := map[string]string{
		"provinceName": "edirne",
		"countyName":   "merkez",
		"townName": "edirne",
	}

	r = mux.SetURLVars(r, vars)

	GetVillagesOfTown(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, mockGetVillagesOfTown())
}

func mockGetVillagesOfProvince() Province {
	var province Province
	f, err := os.Open("mock-1.json")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	err = json.Unmarshal(byteValue, &province)
	return province
}

func mockGetVillagesOfCounty() County {
	var county County
	f, err := os.Open("mock-2.json")

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

func mockGetVillagesOfTown() Town {
	var town Town
	f, err := os.Open("mock-3.json")

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("err is %v", err)
	}

	err = json.Unmarshal(byteValue, &town)
	return town
}
