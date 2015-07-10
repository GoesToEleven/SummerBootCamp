package main

import (
	"fmt"
	"crypto/md5"
	"os"
	"log"
	"io"
	"io/ioutil"
)

func main() {

	// open file
	f, err := os.Open("../resources/table.csv")
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	fmt.Printf("The hash (sum) is: %x", h.Sum(nil))

	// or this way


}