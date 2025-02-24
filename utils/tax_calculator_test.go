package utils

import (
	"itc/models"
	"testing"
)

func TestCalculateTax_NewRegime(t *testing.T) {
	req := models.TaxRequest{
		GrossIncome:    1000000,
		EmployerPF:     50000,
		EmployeePF:     50000,
		NPSContribution: 20000,
		OptOldRegime:   false,
	}

	expectedTaxableIncome := 880000 // 1000000 - 50000 - 50000 - 20000
	result := CalculateTax(req)

	if result.TaxableIncome != expectedTaxableIncome {
		t.Errorf("Expected taxable income %v, got %v", expectedTaxableIncome, result.TaxableIncome)
	}
}

func TestCalculateTax_OldRegime_WithCapitalGains(t *testing.T) {
	req := models.TaxRequest{
		GrossIncome:    1200000,
		OptOldRegime:   true,
		CapitalGains:   true,
		ShortTermGains: 100000,
		LongTermGains:  200000,
	}

	expectedTax := 100000*0.20 + (200000-125000)*0.125 // Capital gain taxes
	result := CalculateTax(req)

	if result.TaxPayable != expectedTax {
		t.Errorf("Expected tax payable %v, got %v", expectedTax, result.TaxPayable)
	}
}

