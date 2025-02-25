package models

type TaxResponse struct {
    TaxableIncome float64 `json:"taxable_income"`
    TaxPayable    float64 `json:"tax_payable"`
    SuggestedITR  string  `json:"suggested_itr"`
}

func NewTaxResponse(taxableIncome, taxPayable float64, suggestedItrType string) *TaxResponse {
    return &TaxResponse {
        TaxableIncome: taxableIncome,
        TaxPayable:    taxPayable,
        SuggestedITR:  suggestedItrType,
    }
}
