package main

import "math"

type CalcRequest struct {
	Value float64 `json:"value" xml:"value" form:"value"`
}

func annualRateFromMonthlyRate(monthlyRate float64) float64 {
	return math.Pow(1+monthlyRate, 12.0) - 1.0
}

func monthlyRateFromAnnualRate(annualRate float64) float64 {
	return math.Pow(1+annualRate, 1.0/12.0) - 1.0
}
