package handlers

import (
    "encoding/json"
    "net/http"
    "itc/models"
    "itc/utils"
)

func CalculateTax(w http.ResponseWriter, r *http.Request) {
    var taxReq models.TaxRequest
    if err := json.NewDecoder(r.Body).Decode(&taxReq); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    taxRes := utils.CalculateTax(taxReq)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(taxRes)
}
