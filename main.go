package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/antoniocascais/torn_stocks/data"
	"github.com/antoniocascais/torn_stocks/strategy"
	"github.com/emirpasic/gods/maps/treemap"
)

var stocksData *treemap.Map

func main() {
	var err error
	stocksData, err = data.LoadData("stocks_2020-02-16.csv")
	checkErr(err)

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/stockInfo", stockHandler)

	http.ListenAndServe(":3000", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Panicf("An error occurred: %v", err)
	}
}

func stockHandler(rw http.ResponseWriter, req *http.Request) {
	params, err := url.ParseQuery(req.URL.RawQuery)
	checkErr(err)

	accronym := params["stock"][0]

	avg, err := strategy.CalculateAverageValue(stocksData, accronym)
	min, dmin := strategy.CalculateMinValue(stocksData, accronym)
	max, dmax := strategy.CalculateMaxValue(stocksData, accronym)

	fmt.Fprintf(rw, "%s average price: %f\n", accronym, avg)
	fmt.Fprintf(rw, "%s minimum price: %f was reached on %s\n", accronym, min, dmin)
	fmt.Fprintf(rw, "%s maximum price: %f was reached on %s\n", accronym, max, dmax)
}
