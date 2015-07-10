package main

import (
	"fmt"
)

func iterate(x chan int) {
	for i := 0; i < 10; i++ {
		x <- i
	}
}

func printer(y chan int) {
	for {
		msg := <-y
		fmt.Println(msg)
	}
}

func main() {
	var c chan int = make(chan int)
	var d chan int = make(chan int)

	go iterate(c)
	go printer(c)
	go iterate(d)
	go printer(d)

	var input string
	fmt.Scanln(&input)
}
