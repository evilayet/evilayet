package county

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCounties(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/provinces", nil)

	GetCounties(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCounty(t *testing.T) {
	t.Parallel()

	var counties []County
	var res []County
	counties = append(counties, County{CountyName: "Enez"})
	counties = append(counties, County{CountyName: "Havsa"})
	counties = append(counties, County{CountyName: "Keşan"})
	counties = append(counties, County{CountyName: "Lalapaşa"})
	counties = append(counties, County{CountyName: "Meriç"})
	counties = append(counties, County{CountyName: "Merkez"})
	counties = append(counties, County{CountyName: "Süloğlu"})
	counties = append(counties, County{CountyName: "Uzunköprü"})
	counties = append(counties, County{CountyName: "İpsala"})

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/counties/{provinceName}", nil)

	vars := map[string]string{
		"provinceName": "edirne",
	}

	r = mux.SetURLVars(r, vars)

	GetCounty(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", counties)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, counties)
}
