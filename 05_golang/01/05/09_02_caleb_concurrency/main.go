package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Task:", n, "Step:", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}