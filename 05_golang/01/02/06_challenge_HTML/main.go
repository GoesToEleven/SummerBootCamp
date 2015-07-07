package main

import (
	"fmt"
	"strconv"
)

const milesToKM float64 = 1.609344

func main() {
	var miles float64
	fmt.Print("Enter miles: ")
	fmt.Scanf("%v", &miles)
	km := miles * milesToKM
	strDivider := "+-------------------------+\n"
	strMiles := "| Miles: " + strconv.FormatFloat(miles, 'f', 0, 64) + "               |\n"
	strKM := "| Kilometers: " + strconv.FormatFloat(km, 'f', 2, 64) + "       |\n"
	fmt.Println("",strDivider, strMiles, strDivider, strKM, strDivider)
}
