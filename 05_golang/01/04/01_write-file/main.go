package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	myStr := "Hello world"

	fmt.Println(myStr)
	fmt.Println([]byte(myStr))

		err := ioutil.WriteFile("hello.txt", []byte(myStr), 0777)
		if err != nil {
			panic("something went wrong")
		}
}
