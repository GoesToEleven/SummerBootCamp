package main

import "fmt"

func main() {
	fmt.Println(vari(7, 8, 9, 10))
}

func vari(args ...int) int {
	var greatest int
	for key, val := range args {
		if key == 0 {
			greatest = val
		}
		if greatest < val {
			greatest = val
		}
	}
	return greatest
}

/*
Write a function with one variadic parameter
that finds the greatest number in a list of numbers.
*/
