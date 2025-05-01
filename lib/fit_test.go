package ris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fittingIsInRange(fit Params, male bool) bool {
	if male {
		if fit.A < 300 && fit.A > 350 {
			return false
		}
		if fit.B < 0 && fit.B > 1 {
			return false
		}
		if fit.K < 500 && fit.K > 550 {
			return false
		}
		if fit.Q < 0 && fit.Q > 20 {
			return false
		}
		if fit.V < 200 && fit.V > 250 {
			return false
		}
		return true
	}

	// female
	if fit.A < 150 && fit.A > 250 {
		return false
	}
	if fit.B < 0 && fit.B > 1 {
		return false
	}
	if fit.K < 200 && fit.K > 300 {
		return false
	}
	if fit.Q < 0 && fit.Q > 20 {
		return false
	}
	if fit.V < 50 && fit.V > 100 {
		return false
	}
	return true
}

func TestFitRISParams(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]bool{
		"male":   true,
		"female": false,
	}

	for test, male := range tests {
		t.Run(test+"-nelder", func(t *testing.T) {
			data, err := LoadDataFromCSV("testdata/" + test + ".csv")
			assert.NoError(err)

			fit, err := FitRISParams(data, 100)
			assert.NoError(err)
			assert.True(fittingIsInRange(fit.Params, male))
		})
	}

	t.Run("test-nelder-total-2023-comparison", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/male.csv")
		assert.NoError(err)

		fit, err := FitRISParams(data, 100)
		assert.NoError(err)
		assert.True(fittingIsInRange(fit.Params, true))

		xavier := RIS(600, 92.45, fit.Params)
		assert.Greater(xavier, 100.0)
		assert.Less(xavier, 150.0)
	})
}
