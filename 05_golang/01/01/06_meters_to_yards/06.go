package main

import "fmt"

const metersToYards float64 = 1.09361

func main(){
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scanf("%v", &meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
