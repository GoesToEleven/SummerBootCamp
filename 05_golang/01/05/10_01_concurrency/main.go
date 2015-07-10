package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		go fmt.Println("Hello world")
	}
	var input string
	fmt.Scanln(&input)
}
