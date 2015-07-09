package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
)

func main() {
	// 0. get csv file - http://statetable.com/
	// 1. open and read csv file
	// 2. parse csv
	// 3. show results

	f, err := os.Open("../resources/state_table.csv")
	if err != nil {
		log.Fatalln("couldn't open csv file", err.Error())
	}

	myReader := csv.NewReader(f)

//	fmt.Println(f)
//	fmt.Println(myReader)

	records, err := myReader.ReadAll()
	if err != nil {
		log.Fatalln("couldn't read it!", err.Error())
	}

	for key, value := range records {
		fmt.Println(key,"-",value)
	}


//	func (r *Reader) ReadAll() (records [][]string, err error)


}
