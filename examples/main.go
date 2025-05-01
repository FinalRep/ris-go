package main

import (
	"fmt"
	"log"

	ris "ris-go/lib"
)

func main() {
	// load data
	maleData, err := ris.LoadDataFromCSV("data/male.csv")
	if err != nil {
		log.Fatal(err)
	}
	femaleData, err := ris.LoadDataFromCSV("data/female.csv")
	if err != nil {
		log.Fatal(err)
	}

	// calculate fitting parameters
	maleFit, err := ris.FitRISParams(maleData, 100)
	if err != nil {
		log.Fatal(err)
	}

	// calculate fitting parameters
	femaleFit, err := ris.FitRISParams(femaleData, 100)
	if err != nil {
		log.Fatal(err)
	}

	// plot result
	if err := ris.PlotFitGraph(maleData, maleFit, "Generalized Logistic Fit Male", "male.png"); err != nil {
		log.Fatal(err)
	}
	// plot result
	if err := ris.PlotFitGraph(femaleData, femaleFit, "Generalized Logistic Fit Female", "female.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fitting Parameters Male: %+v\n", maleFit)
	fmt.Printf("Fitting Parameters Female: %+v\n", femaleFit)

	// check against real values for streetlifting
	// Xavier achieved the first 600 total in 2023
	// Bodyweight: 92.45, Total: 600, 2023 RIS: 115.74819176961559
	xavier := ris.RIS(600.0, 92.45, maleFit.Params)
	fmt.Printf("First 600 Xavier 2023 expected: 115.75, calculated: %.2f\n", xavier)
}
