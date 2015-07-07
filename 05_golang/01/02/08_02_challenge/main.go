package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const miTokm float64 = 1.60934

func main() {
	myFrom := os.Args[1]
	myTo := os.Args[2]

	switch {
	case strings.HasSuffix(myFrom, "mi"):
		switch myTo {
		case "km":
			fmt.Println(myFrom[:len(myFrom)-2])
			fmt.Println(myFrom[0 : len(myFrom)-2])
			miles, _ := strconv.ParseFloat(myFrom[:len(myFrom)-2], 64)
			fmt.Println(miles * miTokm)
		}
	}
}

/*

at terminal:
go install
08_02_challenge 4000mi km

returns this:
4000
4000
6437.36
*/