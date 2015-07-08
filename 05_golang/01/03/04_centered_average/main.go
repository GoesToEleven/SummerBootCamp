package main

import (
	"fmt"
"sort"
)

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(x)
	sort.Ints(x)
	fmt.Println(x)
	y := x[1:]
	fmt.Println(y)
	z := y[:(len(y)-1)]
	fmt.Println(z)

	total := 0

	for _, value := range z {
		total += value
	}

	fmt.Println(total)
	fmt.Println(total/len(z))
}

/*
Implement a centeredAverage function
that computes the average of a list of numbers,
but removes the largest and smallest values.

centeredAverage([]float64{1, 2, 3, 4, 100}) → 3
*/
