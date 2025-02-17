package main

type ShoppingRequest struct {
	BaseValue      float64 `json:"base_value" xml:"base_value" form:"base_value"`
	DurationMonths int     `json:"duration_months" xml:"duration_months" form:"duration_months"`
	InvestmentRate float64 `json:"investment_rate" xml:"investment_rate" form:"investment_rate"`
	DiscountPct    float64 `json:"discount_pct" xml:"discount_pct" form:"discount_pct"`
	ApplyTax       bool    `json:"apply_tax" xml:"apply_tax" form:"apply_tax"`
}

func simulateInvestmentWithPayments(baseValue float64, durationMonths int, rate float64, applyTax bool) (float64, error) {
	total := baseValue
	for i := 0; i < durationMonths; i++ {
		total *= 1.0 + rate
		total -= baseValue / float64(durationMonths)
	}
	if applyTax {
		taxedTotal, err := applyTaxes(total, durationMonths)
		if err != nil {
			return 0.0, err
		}
		return taxedTotal, nil
	} else {
		return total, nil
	}
}
