package parse

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

// ../../../../resources/table.csv

func Parse(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	rows, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	records := make([]Record, 0, len(rows))

	for _, row := range rows {
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records

}
