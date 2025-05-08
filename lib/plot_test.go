package ris

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlotFitGraph(t *testing.T) {
	assert := assert.New(t)

	t.Run("plot", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/male.csv")
		assert.NoError(err)

		fit, err := FitRISParams(data, 100)
		assert.NoError(err)

		err = PlotFitGraph(data, fit, "test", "test.png")
		assert.NoError(err)

		defer os.Remove("test.png")
	})
}
