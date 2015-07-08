package main

import "fmt"

func main() {

	// doesn't work

//	if 0 {
//		fmt.Println("0")
//	}
//
//	if 1 {
//		fmt.Println("1")
//	}
//
//	if 23 {
//		fmt.Println("23")
//	}

	// does work
	remainder := 1
	if remainder != 0 {
		fmt.Println("!= 0")
	}

}
