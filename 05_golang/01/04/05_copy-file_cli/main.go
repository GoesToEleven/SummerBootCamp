package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	toCopy := os.Args[1]
	newCopy := os.Args[2]

	bs, err := ioutil.ReadFile(toCopy)
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}

	f, err2 := os.Create(newCopy)
	if err2 != nil {
		log.Fatalln("couldn't create file", err2.Error())
	}
	defer f.Close()

	_, err4 := f.Write(bs)
	if err4 != nil {
		log.Fatalln("couldn't write byte slice", err4.Error())
	}

	fmt.Println("Files copied")
}
