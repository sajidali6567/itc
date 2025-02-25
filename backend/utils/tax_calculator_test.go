package utils

import (
	"itc/models"
	"testing"
)

// Test cases for CalculateTax Function

func TestCalculateTax_NewRegime(t *testing.T) {
	req := models.TaxRequest{
		GrossIncome:    1000000,
		EmployerPF:     50000,
		EmployeePF:     50000,
		EmployerNPSContribution: 20000,
		OptOldRegime:   false,
	}

	expectedTaxableIncome := 805000.0 // 1000000 - 50000 - 50000 - 20000 - 75000 (S.D)
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


// Test cases for CalculateTaxForNewRegime Function
func TestCalculateTaxForNewRegime(t *testing.T) {
	tests := []struct {
		name           string
		taxableIncome  float64
		expectedTax    float64
	}{
		{"Income 3L (No tax)", 300000, 0},
		{"Income 4L (No tax)", 400000, 0},
		{"Income 5L (5% on 1L)", 500000, 5000},
		{"Income 8L (5% on 4L)", 800000, 20000},
		{"Income 9L (5% on 4L + 10% on 1L)", 900000, 30000},
		{"Income 12L (5% on 4L + 10% on 4L)", 1200000, 60000},
		{"Income 14L (5% on 4L + 10% on 4L + 15% on 2L)", 1400000, 90000},
		{"Income 16L (5% on 4L + 10% on 4L + 15% on 4L)", 1600000, 120000},
		{"Income 18L (5% on 4L + 10% on 4L + 15% on 4L + 20% on 2L)", 1800000, 160000},
		{"Income 20L (5% on 4L + 10% on 4L + 15% on 4L + 20% on 4L)", 2000000, 200000},
		{"Income 22L (5% on 4L + 10% on 4L + 15% on 4L + 20% on 4L + 25% on 2L)", 2200000, 250000},
		{"Income 24L (5% on 4L + 10% on 4L + 15% on 4L + 20% on 4L + 25% on 4L)", 2400000, 300000},
		{"Income 26L (5% on 4L + 10% on 4L + 15% on 4L + 20% on 4L + 25% on 4L + 30% on 2L)", 2600000, 360000},
		{"Income 30L (5% on 4L + 10% on 4L + 15% on 4L + 20% on 4L + 25% on 4L + 30% on 6L)", 3000000, 480000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateTaxForNewRegime(tt.taxableIncome)
			if got != tt.expectedTax {
				t.Errorf("calculateTaxForNewRegime(%f) = %f; want %f", tt.taxableIncome, got, tt.expectedTax)
			}
		})
	}
}
