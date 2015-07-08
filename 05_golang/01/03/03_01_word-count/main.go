package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "whatever I want it to be"
	for key, value := range str {
		fmt.Println(key, value, string(value))
	}

	fmt.Println(strings.Split(str, " "))
	fmt.Printf("%q\n", strings.Split(str, " "))

	var words []string = strings.Split(str, " ")
	fmt.Println(words)

	for k, v := range words {
		fmt.Println(k, v)
	}
}
