package main

import (
	"log"
	"os"
	"io"
	"fmt"
)

func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	for {
		bs := make([]byte, 1)
		n, err := srcFile.Read(bs)
		if err == io.EOF {
			break
		} else if err != nil {
			// do something with error
			log.Fatalln(err)
		}
		if n == 0 {
			break
		}

		actualData := bs[:n]
		log.Println(actualData)

	}
}