package main

import (
	"fmt"
	"log"

	rislib "ris-go/lib"
)

func main() {
	data, err := rislib.LoadDataFromCSV("lib/testdata/test.csv")
	if err != nil {
		log.Fatal(err)
	}

	initial := rislib.RISParams{A: 10, K: 100, Q: 1, B: 0.05, V: 75}
	fitted, err := rislib.FitRISParams(data, initial)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Optimierte Parameter: %+v\n", fitted)
	for _, dp := range data {
		score := rislib.RIS(dp.Total, dp.Bodyweight, fitted)
		fmt.Printf("RIS für %.1f kg BW / %.1f kg Total: %.2f\n", dp.Bodyweight, dp.Total, score)
	}
}
