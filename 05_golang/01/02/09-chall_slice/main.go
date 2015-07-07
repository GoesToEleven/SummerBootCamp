package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var smallest, largest int

	for key, value := range x {
		if key == 0 {
			smallest = value
			largest = value
		}
		if smallest < value {
			smallest = value
		}
		if largest > value {
			largest = value
		}
		fmt.Println(key, value)
	}
	fmt.Println(`smallest:\t`,smallest)
	fmt.Println(`largest:\t`,largest)
	fmt.Println("smallest:\t",smallest)
	fmt.Println("largest:\t",largest)
}
