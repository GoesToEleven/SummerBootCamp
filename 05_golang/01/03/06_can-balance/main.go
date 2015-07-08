package main

import "fmt"

func canBalance(y []int) bool {
	var total int

	for _, val := range y {
		total += val
	}

	remainder := total % 2
	fmt.Println(remainder)

	if remainder != 0 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(canBalance([]int{1, 1, 1, 2, 1}))
	fmt.Println(canBalance([]int{2, 1, 1, 2, 1}))
	fmt.Println(canBalance([]int{10, 10}))
}


/*
Given a non-empty list,
return true if there is a place to split the list
so that the sum of the numbers on one side is equal to
the sum of the numbers on the other side.

	canBalance([]int{1, 1, 1, 2, 1}) → true
	canBalance([]int{2, 1, 1, 2, 1}) → false
	canBalance([]int{10, 10}) → true
*/