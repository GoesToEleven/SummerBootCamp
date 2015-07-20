package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type ToDo struct {
	Item string
	Id   string
}

func main() {
	fmt.Println(ToDo{})

	fmt.Println(ToDo{
		Item: "first item",
		Id:   "1",
	})

	second := ToDo{
		"second",
		"2",
	}
	fmt.Println(second)

	bs, err := json.Marshal(second)
	if err != nil {
		log.Fatalln("error", err)
	}
	fmt.Println(string(bs))

	third := ToDo{
		Item: "third",
	}
	fmt.Println(third)

	third.Id = "33"

	bs, err = json.Marshal(third)
	if err != nil {
		log.Fatalln("error", err)
	}
	fmt.Println(string(bs))

}
