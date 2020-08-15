package province

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProvinces(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/provinces", nil)

	GetProvinces(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProvince(t *testing.T) {
	t.Parallel()

	mock := Province{ProvinceName: "Edirne"}
	var res Province

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r, _ := http.NewRequest("GET", "/provinces/{provinceName}", nil)

	vars := map[string]string{
		"provinceName": "edi",
	}

	r = mux.SetURLVars(r, vars)

	GetProvince(w, r)

	err := json.Unmarshal(w.Body.Bytes(), &res)

	if err != nil {
		t.Logf("Cant deserialize %v", err)
	}

	t.Logf("%v", res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mock, res)
}
