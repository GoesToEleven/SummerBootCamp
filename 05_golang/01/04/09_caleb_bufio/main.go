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


// might not be a complete working file
// this file was just pulled part way through lecture
// https://www.youtube.com/playlist?list=PLSak_q1UXfPqJ6pIpJt2oNJTPx9ByBy6K