package utils

import "itc/models"

const (
    NEW_REGIME_STANDARD_DEDUCTION = 75000
    OLD_REGIME_STANDARD_DEDUCTION = 50000
)

func CalculateTax(req models.TaxRequest) models.TaxResponse {
    var tax float64
    var itrType string

    // TO DO: Limit NPS contribution based on income tax and 1.5L
    // Standard deduction needs to be taken into account for both regime
    taxableIncome := req.GrossIncome - req.EmployerPF - req.EmployeePF - req.EmployerNPSContribution

    // if employee opts for new regime, calculate the tax and return
    if req.OptOldRegime == false {

        taxableIncome -= NEW_REGIME_STANDARD_DEDUCTION

        tax = calculateTaxForNewRegime(taxableIncome)

        // To DO: Use NewTaxResponse() function 
        taxRes := models.TaxResponse{
            TaxableIncome: taxableIncome,
            TaxPayable:    tax,
            SuggestedITR:  itrType,
        }
    
        return taxRes
    }

    taxableIncome -= OLD_REGIME_STANDARD_DEDUCTION

    // Old regime flow
    tax = calculateTaxForOldRegime(&req, taxableIncome)

    taxRes := models.TaxResponse{
        TaxableIncome: taxableIncome,
        TaxPayable:    tax,
        SuggestedITR:  itrType,
    }

    return taxRes
}


/* 
 * Input: Taxable Income (Gross Income - All applicable deductions)
 * Output: Tax Payable
 * ----------------------------------------------------------------
 * This functions calculates tax as per following slabs
 * Slab is applicable from FY 24-25 onwards
 * 0-4L - nil
 * 4-8L - 5%
 * 8-12L - 10%
 * 12-16L - 15%
 * 16-20L - 20%
 * 20-24L - 25%
 * 24L above - 30%
 */

 func calculateTaxForNewRegime(taxableIncome float64) float64 {
    var tax float64

    slabs := []struct {
        limit float64
        rate  float64
    }{
        {400000, 0.00},
        {800000, 0.05},
        {1200000, 0.10},
        {1600000, 0.15},
        {2000000, 0.20},
        {2400000, 0.25},
    }

    prevLimit := 0.0
    for _, slab := range slabs {
        if taxableIncome > slab.limit {
            tax += (slab.limit - prevLimit) * slab.rate
            prevLimit = slab.limit
        } else {
            tax += (taxableIncome - prevLimit) * slab.rate
            return tax
        }
    }

    // If income exceeds 24L, apply 30% on the remaining amount
    if taxableIncome > 2400000 {
        tax += (taxableIncome - 2400000) * 0.30
    }

    return tax
}

func calculateTaxForOldRegime(req *models.TaxRequest, taxableIncome float64) float64 {
    var tax float64

    taxableIncome -= req.SavingAccountInterest - 10000
    taxableIncome -= req.Section80CInvest
    taxableIncome -= req.Section80DHealth
    

    if req.CapitalGains {
        tax += req.ShortTermGains * 0.20
        if req.LongTermGains > 125000 {
            tax += (req.LongTermGains - 125000) * 0.125
        }
        // itrType = "ITR-2"
    } else {
        // itrType = "ITR-1"
    }
    
    return tax
}