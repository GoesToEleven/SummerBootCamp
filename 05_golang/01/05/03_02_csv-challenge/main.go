package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type oneRecord struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
	adjClose float64
}

func main() {
	// 1. Get csv file
	// 2. open file
	// 3. read file
	// 4. parse file
	// 5. display file

	// 2. open file
	src, err := os.Open("../resources/table.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	// src is a *File
	// func (f *File) Read(b []byte) (n int, err error)
	// func NewReader(r io.Reader) *Reader

	// 3. read file
	rdr := csv.NewReader(src)
	// rdr is a *Reader
	// func (r *Reader) ReadAll() (records [][]string, err error)

	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	// 4. parse file
	myData := make(map[int]oneRecord)

	for k, v := range data {
		fmt.Println(k, v)
		for _, value := range v {
			myData[k].Date = value
		}
	}

}

/*
Grab historical financial data from Yahoo
http://finance.yahoo.com/q/hp?s=GOOG+Historical+Prices
as a csv file, read that file and dumps the contents as an HTML table.
EXTRA CHALLENGE: draw a line chart of the data.
*/
