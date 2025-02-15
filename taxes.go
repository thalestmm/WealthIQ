package main

import "errors"

// OBS: All values are stored in DAYS
const (
	IrLess180  float64 = 0.225
	Ir181To360 float64 = 0.2
	Ir361To720 float64 = 0.175
	IrOver720  float64 = 0.15
)

func applyTaxes(baseValue float64, durationMonths int) (float64, error) {
	if durationMonths <= 0 {
		return 0.0, errors.New("duration must be greater than zero")
	} else if durationMonths <= 6 {
		return (1.0 - IrLess180) * baseValue, nil
	} else if durationMonths <= 12 {
		return (1.0 - Ir181To360) * baseValue, nil
	} else if durationMonths <= 24 {
		return (1.0 - Ir361To720) * baseValue, nil
	} else {
		return (1.0 - IrOver720) * baseValue, nil
	}

}
