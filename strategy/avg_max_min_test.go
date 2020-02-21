package strategy

import (
	"testing"

	"github.com/antoniocascais/torn_stocks/data"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	d, _ := data.LoadData("../test/sample_data.csv")

	avg, err := CalculateAverageValue(d, "BAG")

	assert.Equal(t, 798.98, avg)
	assert.NoError(t, err)
}

func TestCalculateMin(t *testing.T) {
	d, _ := data.LoadData("../test/sample_data.csv")

	min, date := CalculateMinValue(d, "BAG")

	assert.Equal(t, 798.84, min)
	assert.Equal(t, "2019-03-21 00:15:05", date)
}

func TestCalculateMax(t *testing.T) {
	d, _ := data.LoadData("../test/sample_data.csv")

	max, date := CalculateMaxValue(d, "SYM")

	assert.Equal(t, 295.854, max)
	assert.Equal(t, "2019-03-21 00:00:08", date)
}
