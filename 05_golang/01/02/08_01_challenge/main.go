package main

import (
	"os"
	"fmt"
)

func main() {
	from := os.Args[1]
	to := os.Args[2]

	fmt.Println(from, to)
}
/*

at terminal:
go install
08_01_challenge 40 50

returns this:
40 50


*/