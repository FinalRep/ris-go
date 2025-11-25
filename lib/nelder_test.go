package ris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFitRISParamsNelder(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]bool{
		"male":   true,
		"female": false,
	}

	for test, male := range tests {
		t.Run(test+"-nelder", func(t *testing.T) {
			data, err := LoadDataFromCSV("testdata/" + test + ".csv")
			assert.NoError(err)

			fit, err := FitRISParamsNelder(data, 100)
			assert.NoError(err)
			assert.True(fittingIsInRange(fit.Params, male))
		})
	}

	t.Run("test-nelder-total-2023-comparison", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/male.csv")
		assert.NoError(err)

		fit, err := FitRISParamsNelder(data, 100)
		assert.NoError(err)
		assert.True(fittingIsInRange(fit.Params, true))

		xavier := RIS(600, 92.45, fit.Params)
		assert.Greater(xavier, 100.0)
		assert.Less(xavier, 150.0)
	})
}
