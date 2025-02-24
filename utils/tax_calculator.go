package utils

import "itc/models"

func CalculateTax(req models.TaxRequest) models.TaxResponse {
    var tax float64
    var itrType string

    taxableIncome := req.GrossIncome - req.EmployerPF - req.EmployeePF - req.NPSContribution
    if req.OptOldRegime {
        taxableIncome -= req.SavingsInterest - 10000
        taxableIncome -= req.Section80CInvest
        taxableIncome -= req.Section80DHealth
    }

    if req.CapitalGains {
        tax += req.ShortTermGains * 0.20
        if req.LongTermGains > 125000 {
            tax += (req.LongTermGains - 125000) * 0.125
        }
        itrType = "ITR-2"
    } else {
        itrType = "ITR-1"
    }

    taxRes := models.TaxResponse{
        TaxableIncome: taxableIncome,
        TaxPayable:    tax,
        SuggestedITR:  itrType,
    }

    return taxRes
}
