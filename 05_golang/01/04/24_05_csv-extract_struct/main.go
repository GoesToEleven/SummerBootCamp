package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// make a struct
// give some "STRUCTURE" to your data
type state struct {
	abbreviation string
	name         string
	ocean        string
}

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

	records, err := myReader.ReadAll()
	if err != nil {
		log.Fatalln("couldn't read it!", err.Error())
	}

	// now add your struct to your map
	myStates := make(map[string]state)

	for _, value := range records {

		newState := state{
			abbreviation: value[2],
			name:         value[1],
			ocean:        value[15],
		}

		myStates[value[2]] = newState
	}

	lookup := os.Args[1]

	fmt.Println(myStates[lookup])

	// run go install
	// run program from command line (cli)
	// add an argument <STATE ABBREVIATION>

}

// example of what is in each record:
//	[47 Washington WA USA state 10 current occupied  53 Wash. X 4 West 9 Pacific 9]
