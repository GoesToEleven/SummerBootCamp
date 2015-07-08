package main

import "fmt"

func rotate(args ...*int) {
	if len(args) == 0 {
		return
	}

	firstVal := *args[0]

	for i := 0; i < len(args)-1; i++ {
		*args[i] = *args[i+1]
	}
	*args[len(args)-1] = firstVal
}

func main() {
	x := 1
	y := 2
	z := 3
	fmt.Println("x", x, &x)
	fmt.Println("y", y, &y)
	fmt.Println("z", z, &z)
	rotate(&x, &y, &z)
	fmt.Println("x", x, &x)
	fmt.Println("y", y, &y)
	fmt.Println("z", z, &z)
	//	m := 2
	//	m = m + 1
}

/*
Write a program that can rotate two integers
(x := 1; y := 2; rotate(&x, &y) should give you x=2 and y=1)

*/
