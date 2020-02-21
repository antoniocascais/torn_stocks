package strategy

import (
	"log"
	"strconv"
	"strings"
)

func CalculateProfits(entries []string) float64 {
	row := strings.Split(entries[0], ",")
	first := row[3]

	row = strings.Split(entries[len(entries)-1], ",")
	last := row[3]

	f, err := strconv.ParseFloat(first, 64)
	checkErr(err)

	l, err := strconv.ParseFloat(last, 64)
	checkErr(err)

	return l - f
}

func checkErr(err error) {
	if err != nil {
		log.Panicf("An error occurred: %v", err)
	}
}
