package ris

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFitRISParamsScipy(t *testing.T) {
	assert := assert.New(t)

	// Check if python3 is available in the environment
	_, err := exec.LookPath("python3")
	if err != nil {
		t.Skip("Skipping Scipy test: python3 not found in PATH")
	}

	tests := map[string]bool{
		"male":   true,
		"female": false,
	}

	for test, male := range tests {
		t.Run(test+"-scipy", func(t *testing.T) {
			data, err := LoadDataFromCSV("testdata/" + test + ".csv")
			assert.NoError(err)

			// Call the new Scipy bridge function
			fit, err := FitRISParamsScipy(data, 100)
			assert.NoError(err)
			
			// Ensure the parameters found by Scipy are physically realistic
			assert.True(fittingIsInRange(fit.Params, male), "Scipy parameters out of expected range for %s", test)
			
			// Verify linear model components were calculated
			assert.NotZero(fit.LineSlope)
			assert.NotZero(fit.LineIntercept)
		})
	}

	t.Run("test-scipy-total-2023-comparison", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/male.csv")
		assert.NoError(err)

		fit, err := FitRISParamsScipy(data, 100)
		assert.NoError(err)

		// Verification with a known high-level performance
		// Xavier (example athlete)
		xavier := RIS(600, 92.45, fit.Params)
		assert.Greater(xavier, 100.0)
		assert.Less(xavier, 150.0)
	})
}