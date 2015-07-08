package main

import "fmt"

func swap(x, y *int) {
	a := *x
	b := *y
	*x = b
	*y = a
}

func main() {
	x := 1
	y := 2
	fmt.Println("x", x, &x)
	fmt.Println("y", y, &y)
	swap(&x,&y)
	fmt.Println("x", x, &x)
	fmt.Println("y", y, &y)
}


/*
Write a program that can swap two integers
(x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1)

*/