package data

import (
	"testing"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/stretchr/testify/assert"
)

func TestLoadData(t *testing.T) {
	res, err := LoadData("../test/sample_data.csv")

	expectedRes := buildExpectedData()
	expectedKeys := []int{1553126408, 1553127305}

	// assert that each key has the expected values
	for _, v := range res.Keys() {
		value, _ := res.Get(v)
		expectedValue, found := expectedRes.Get(v)
		assert.True(t, found)
		assert.Equal(t, expectedValue, value)
	}
	// assert that the keys are inserted in the right order
	assert.Equal(t, expectedKeys, castToIntArray(res.Keys()))
	// assert that we have inserted the expected number of keys
	assert.Equal(t, 2, res.Size())
	// assert there are no errors
	assert.NoError(t, err)
}

func buildExpectedData() *treemap.Map {
	m := treemap.NewWithIntComparator()

	data1 := make([]StockData, 0)
	data1 = append(data1, StockData{Accronym: "BAG", Date: "2019-03-21 00:00:08", Price: 799.12})
	data1 = append(data1, StockData{Accronym: "WLT", Date: "2019-03-21 00:00:08", Price: 497.319})
	data1 = append(data1, StockData{Accronym: "SYM", Date: "2019-03-21 00:00:08", Price: 295.854})
	m.Put(1553126408, data1)

	data2 := make([]StockData, 0)
	data2 = append(data2, StockData{Accronym: "BAG", Date: "2019-03-21 00:15:05", Price: 798.84})
	data2 = append(data2, StockData{Accronym: "WLT", Date: "2019-03-21 00:15:05", Price: 496.992})
	data2 = append(data2, StockData{Accronym: "SYM", Date: "2019-03-21 00:15:05", Price: 294.819})
	m.Put(1553127305, data2)

	return m
}

func castToIntArray(input []interface{}) []int {
	res := make([]int, 0)

	for _, v := range input {
		res = append(res, v.(int))
	}

	return res
}
