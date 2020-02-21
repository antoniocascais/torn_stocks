package strategy

import (
	"math"

	"github.com/antoniocascais/torn_stocks/data"
	"github.com/emirpasic/gods/maps/treemap"
)

func CalculateAverageValue(stocksData *treemap.Map, accronym string) (float64, error) {
	stockEntries := 0
	stockValue := 0.0

	for _, k := range stocksData.Keys() {
		td, _ := stocksData.Get(k)
		var timestampData = td.([]data.StockData)

		for i, _ := range timestampData {
			if timestampData[i].Accronym == accronym {
				stockEntries = stockEntries + 1
				stockValue = stockValue + timestampData[i].Price
			}
		}
	}
	return stockValue / float64(stockEntries), nil
}

func CalculateMinValue(stocksData *treemap.Map, accronym string) (float64, string) {
	n := math.MaxFloat64
	var min = &n
	var sd = &data.StockData{}
	d:= ""
	var date = &d

	calculateValue(stocksData, accronym, func(p float64, d string) {
		if *min > p {
			*min = p
			*date = d
		}
		sd = &data.StockData{Price: *min, Date: *date}
	})
	return sd.Price, sd.Date
}

func CalculateMaxValue(stocksData *treemap.Map, accronym string) (float64, string) {
	n := 0.0
	var max = &n
	var sd = &data.StockData{}
	d:= ""
	var date = &d

	calculateValue(stocksData, accronym, func(p float64, d string) {
		if *max < p {
			*max = p
			*date = d
		}
		sd = &data.StockData{Price: *max, Date: *date}
	})

	return sd.Price, sd.Date
}

func calculateValue(stocksData *treemap.Map, accronym string, alg func(float64, string)) {
	for _, k := range stocksData.Keys() {
		td, _ := stocksData.Get(k)
		var timestampData = td.([]data.StockData)

		for i, _ := range timestampData {
			if timestampData[i].Accronym == accronym {
				alg(timestampData[i].Price, timestampData[i].Date)
			}
		}
	}
}
