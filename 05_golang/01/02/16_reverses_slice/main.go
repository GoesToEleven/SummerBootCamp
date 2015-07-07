package main

import "fmt"

func main() {
	zs := [1,2,3,4,5,]
	reverse()
}

func reverse(xs []int) []int {
	ys := make([]int, (len(xs)-1))
	for k, v := range xs {
		ys[len(xs)-(k+1)] = v
	}
	return ys
}
