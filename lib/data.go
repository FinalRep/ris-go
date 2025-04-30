package rislib

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
	for _, row := range rows[1:] { // Überspringt Header
		bw, _ := strconv.ParseFloat(row[0], 64)
		total, _ := strconv.ParseFloat(row[1], 64)
		data = append(data, DataPoint{
			Bodyweight: bw,
			Total:      total,
		})
	}
	return data, nil
}
