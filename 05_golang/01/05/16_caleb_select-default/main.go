package main

import (
	"fmt"
)

func writer(msg string, c chan string) {
	for {
		c <- msg
	}
}

func main() {
	c1, c2 := make(chan string), make(chan string)

	go writer("Writer #1", c1)
	go writer("Writer #2", c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("GOT MESSAGE 1", msg1)
		case msg2 := <-c2:
			fmt.Println("GOT MESSAGE 2", msg2)
		default:
			fmt.Println("Default")
		}
	}

	var input string
	fmt.Scanln(&input)
}
