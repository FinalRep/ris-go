package ris

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
)

// LoadDataFromCSV imports data from a csv file
func LoadDataFromCSV(path string) ([]DataPoint, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := []DataPoint{}
	for _, row := range rows[1:] {
		bw, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, err
		}
		total, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}
		data = append(data, DataPoint{
			BodyWeight: bw,
			Total:      total,
		})
	}
	return data, nil
}
