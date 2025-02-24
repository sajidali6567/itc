package handlers

import (
	"bytes"
	"encoding/json"
	"itc/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateTaxHandler(t *testing.T) {
	reqBody := models.TaxRequest{
		GrossIncome:  1000000,
		OptOldRegime: false,
	}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/calculate-tax", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateTax)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

