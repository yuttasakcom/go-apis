package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yuttasakcom/go-apis/routes"
)

func TestRouterHealthWithSuccess(t *testing.T) {
	r := routes.Router()
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /health is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}
