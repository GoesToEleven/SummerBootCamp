package main

import "fmt"

const mToKm = 1.60934

func main(){
	var miles float64
	fmt.Print("Enter miles: ")
	fmt.Scanf("%f", &miles)
	kilometers := miles * mToKm
    fmt.Println(miles, " miles equals ", kilometers, " kilometers.")
}
