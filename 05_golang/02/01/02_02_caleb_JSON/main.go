package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
	{
	"name": "Todd McLeod",
	"age": 44
	}
	`

	var obj map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
