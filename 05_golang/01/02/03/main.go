package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%.2f", 1.23)
	fmt.Println()
	// creates strings Sprintf
	myStr := fmt.Sprintf("%.2f", 1.33)
	fmt.Println(myStr)
}
