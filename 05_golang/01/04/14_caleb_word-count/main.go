package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(str) {
		counts[word]++
	}
	return counts
}

func main() {
	str := "I have a cat cat cat cat"
	result := WordCount(str)
	fmt.Println("Number of cats:", result["cat"])
}


// might not be a complete working file
// this file was just pulled part way through lecture
// https://www.youtube.com/playlist?list=PLSak_q1UXfPqJ6pIpJt2oNJTPx9ByBy6K
