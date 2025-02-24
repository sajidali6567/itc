package models

type TaxResponse struct {
    TaxableIncome float64 `json:"taxable_income"`
    TaxPayable    float64 `json:"tax_payable"`
    SuggestedITR  string  `json:"suggested_itr"`
}
