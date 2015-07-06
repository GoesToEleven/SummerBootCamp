package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

/*

Write a program that prints out all the numbers
evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)

*/
