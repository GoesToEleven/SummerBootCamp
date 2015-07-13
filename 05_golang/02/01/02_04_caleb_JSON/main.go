package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
	[100, 200, 300.5, 400.1234]
	`

	var obj []float64

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
