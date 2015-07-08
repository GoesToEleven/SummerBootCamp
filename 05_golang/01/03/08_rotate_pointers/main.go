package main

import "fmt"

func rotate(x, y, z *int) {
	*x, *y, *z = *y, *z, *x
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
