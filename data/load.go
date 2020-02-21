package data

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
)

type StockData struct {
	Accronym string
	Price    float64
	Date     string
}

func LoadData(file string) (*treemap.Map, error) {
	var timestampData []StockData

	csvfile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvfile)
	r.Comma = ';'

	stocksDb := treemap.NewWithIntComparator()
	stocksDb.Size()

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		row := strings.Split(record[0], ",")
		accronym := row[1]
		timestamp, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, err
		}
		date := row[4]
		stockPrice, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, err
		}

		tmpData, found := stocksDb.Get(timestamp)

		if !found {
			timestampData = make([]StockData, 0)
		}

		if tmpData != nil {
			timestampData = tmpData.([]StockData)
		}

		timestampData = append(timestampData, StockData{
			Accronym: accronym,
			Price:    stockPrice,
			Date:     date,
		},
		)

		stocksDb.Put(timestamp, timestampData)
	}

	return stocksDb, nil
}
