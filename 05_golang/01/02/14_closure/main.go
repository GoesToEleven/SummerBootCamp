package main

import "fmt"

func main() {
	ys := 0

	myfunc := func() int {
		ys++
		return ys
	}

	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())

}
