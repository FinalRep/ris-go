package main

import (
	"fmt"
	"log"

	ris "ris-go/lib"
)

func main() {
	data, err := ris.LoadDataFromCSV("data/male.csv")
	if err != nil {
		log.Fatal(err)
	}

	fitted, err := ris.FitRISParams(data, 100, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Optimierte Parameter: %+v\n", fitted)
	for _, dp := range data {
		score := ris.RIS(dp.Total, dp.BodyWeight, fitted.Params)
		fmt.Printf("RIS: %.2f, Bodyweight: %.2f, Total: %.2f\n", score, dp.BodyWeight, dp.Total)
	}

	// Xavier first 600 total in 2023
	// Bodyweight: 92.45, Total: 600, 2023 RIS: 115.74819176961559
	xavier := ris.RIS(600.0, 92.45, fitted.Params)
	fmt.Printf("First 600 Xavier 2023 expected: 115.75, calculated: %.2f\n", xavier)
}
