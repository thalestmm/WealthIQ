package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type CDBCompRequest struct {
	BaseCDBRate    float64 `json:"cdb_value" form:"cdb_value" xml:"cdb_value"`
	LCIRate        float64 `json:"lci_value" form:"lci_value" xml:"lci_value"`
	DurationMonths int     `json:"duration_months" form:"duration_months" xml:"duration_months"`
}

func realCdbRate(originalRate float64, durationMonths int) (float64, error) {
	total, err := applyTaxes(originalRate, durationMonths)
	if err != nil {
		return 0.0, err
	}
	return total, nil
}

func compareLciCdb(cdbEndDate time.Time, lciRate float64, cdbRate float64) (string, float64, error) {
	timeDeltaNs := cdbEndDate.Sub(time.Now())
	timeDeltaDays := math.Round(timeDeltaNs.Seconds() / 86400)
	fmt.Printf("Time Delta: %v days", timeDeltaDays)

	newCdbRate, err := realCdbRate(cdbRate, int(timeDeltaDays))
	if err != nil {
		log.Fatal(err)
		return "", 0.0, err
	}

	fmt.Printf("LCI/LCA Rate: \t \t %.2f %s \n", lciRate, "%")
	fmt.Printf("Original CDB Rate: \t %.2f %s \n", cdbRate, "%")
	fmt.Printf("Real CDB Rate: \t \t %.2f %s \n", newCdbRate, "%")

	if lciRate > newCdbRate {
		percDiff := lciRate / newCdbRate
		return "LCI/LCA", percDiff, nil
	} else {
		percDiff := newCdbRate / lciRate
		return "CDB", percDiff, nil
	}

}
