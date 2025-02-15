package main

func simulateInvestment(baseValue float64, durationMonths int, monthlyRate float64, applyTax bool) (float64, error) {
	total := baseValue
	for i := 0; i < durationMonths; i++ {
		total *= 1.0 + monthlyRate
	}
	if applyTax {
		// Apply taxes only over true earnings
		earnings := total - baseValue
		taxedEarnings, err := applyTaxes(earnings, durationMonths)
		taxedTotal := baseValue + taxedEarnings
		if err != nil {
			return 0.0, err
		}
		return taxedTotal, nil
	}
	return total, nil
}

type InvestmentRequest struct {
	BaseValue      float64 `json:"base_value" form:"base_value" xml:"base_value"`
	DurationMonths int     `json:"duration_months" form:"duration_months" xml:"duration_months"`
	MonthlyRate    float64 `json:"monthly_rate" form:"monthly_rate" xml:"monthly_rate"`
	ApplyTax       bool    `json:"apply_taxes" form:"apply_taxes" xml:"apply_taxes"`
}

type InvestmentResponse struct {
	BaseValue      float64
	DurationMonths int
	MonthlyRate    float64
	ApplyTax       bool
	TotalValue     float64
}
