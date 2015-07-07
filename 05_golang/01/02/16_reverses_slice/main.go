package main

import "fmt"

func main() {
	zs := []int{1, 2, 3, 4, 5}
	fmt.Println(reverse(zs))
}

func reverse(xs []int) []int {
	ys := make([]int, len(xs))
	for k, v := range xs {
		ys[len(xs)-(k+1)] = v
	}
	return ys
}
