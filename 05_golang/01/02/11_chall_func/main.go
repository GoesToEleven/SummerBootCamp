package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		x, y := halfer(i)
		fmt.Println(i, i/2, x, y)
	}
}

func halfer(x int) (int, bool) {
	return x / 2, x%2 == 0
}

/*
Write a function which takes an integer
and halves it and returns true if it was even or false if it was odd.
For example half(1) should return (0, false) and half(2) should return (1, true).

*/
