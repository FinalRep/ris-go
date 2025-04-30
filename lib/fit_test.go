package rislib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFitRISParams(t *testing.T) {
	assert := assert.New(t)
	t.Run("test", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/test.csv")
		assert.NoError(err)

		_, err = FitRISParams(data, RISParams{A: 10, K: 100, Q: 1, B: 0.05, V: 75})
		assert.NoError(err)
	})
}
