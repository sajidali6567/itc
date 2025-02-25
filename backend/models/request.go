package models

type TaxRequest struct {
    GrossIncome          float64 `json:"gross_income"`
    EmployerPF          float64 `json:"employer_pf"`
    EmployeePF          float64 `json:"employee_pf"`
    EmployerNPSContribution     float64 `json:"employer_nps_contribution"`
    OptOldRegime        bool    `json:"opt_old_regime"`
    SavingAccountInterest     float64 `json:"saving_account_interest"`
    CapitalGains        bool    `json:"capital_gains"`
    ShortTermGains      float64 `json:"short_term_gains"`
    LongTermGains       float64 `json:"long_term_gains"`
    HRA                 float64 `json:"hra"`
    RentPaid            float64 `json:"rent_paid"`
    MetroCity           bool    `json:"metro_city"`
    Section80CInvest    float64 `json:"section_80c_invest"`
    Section80DHealth    float64 `json:"section_80d_health"`
    SeniorCitizen       bool    `json:"senior_citizen"`
    ParentsSeniorCitizen bool   `json:"parents_senior_citizen"`
    EmployeeNPSContribution     float64 `json:"employee_nps_contribution"`
}
