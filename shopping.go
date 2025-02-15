package main

import (
	"fmt"
	"log"
)

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

func compareDiscountToInvestment(baseValue float64, discountPct float64,
	durationMonths int, monthlyRate float64, applyTax bool) {
	investmentTotal, err := simulateInvestmentWithPayments(baseValue, durationMonths, monthlyRate, applyTax)
	if err != nil {
		log.Fatal(err)
	}
	discountedValue := (1 - discountPct/100.0) * baseValue
	diff := baseValue - discountedValue
	if investmentTotal > diff {
		fmt.Printf("It's better to invest by $%.2f \n", investmentTotal-diff)
	} else {
		fmt.Printf("Buy it with the discount, you'll save %.2f \n", diff-investmentTotal)
	}
}
