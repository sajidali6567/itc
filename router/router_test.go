package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterSetup(t *testing.T) {
	r := SetupRouter()
	req, _ := http.NewRequest("POST", "/calculate-tax", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest { // Expecting BadRequest due to missing JSON
		t.Errorf("Expected BadRequest, got %v", rr.Code)
	}
}

