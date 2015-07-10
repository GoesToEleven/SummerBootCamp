package main

import (
	"os"
	"encoding/csv"
	"log"
	"fmt"
)

func main() {
	// 1. Get csv file
	// 2. open file
	// 3. read file
	// 4. parse file
	// 5. display file

	src, err := os.Open("table.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	// src is a *File
	// func (f *File) Read(b []byte) (n int, err error)
	// func NewReader(r io.Reader) *Reader

	rdr := csv.NewReader(src)
	// rdr is a *Reader
	// func (r *Reader) ReadAll() (records [][]string, err error)

	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	for v, k := range data {
		fmt.Println(v, k)
	}

}




/*
Grab historical financial data from Yahoo
http://finance.yahoo.com/q/hp?s=GOOG+Historical+Prices
as a csv file, read that file and dumps the contents as an HTML table.
EXTRA CHALLENGE: draw a line chart of the data.
 */
