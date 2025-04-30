package main

import (
	"fmt"
	"log"

	ris "ris-go/lib"
)

func main() {
	// load data
	data, err := ris.LoadDataFromCSV("data/male.csv")
	if err != nil {
		log.Fatal(err)
	}

	// calculate fitting parameters
	fitted, err := ris.FitRISParams(data, 100, true)
	if err != nil {
		log.Fatal(err)
	}

	// plot result
	if err := ris.PlotFitGraph(data, fitted, "test.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fitting Parameters: %+v\n", fitted)

	// check against real values for streetlifting
	// Xavier achieved the first 600 total in 2023
	// Bodyweight: 92.45, Total: 600, 2023 RIS: 115.74819176961559
	xavier := ris.RIS(600.0, 92.45, fitted.Params)
	fmt.Printf("First 600 Xavier 2023 expected: 115.75, calculated: %.2f\n", xavier)
}
