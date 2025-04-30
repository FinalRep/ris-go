package rislib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDataFromCSV(t *testing.T) {
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/test.csv")
		assert.NoError(err)
		assert.NotNil(data)
	})
	t.Run("not-found", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/not-found.csv")
		assert.Error(err)
		assert.Nil(data)
	})
	t.Run("broken-csv", func(t *testing.T) {
		data, err := LoadDataFromCSV("testdata/broken.csv")
		assert.Error(err)
		assert.Nil(data)
	})
}
