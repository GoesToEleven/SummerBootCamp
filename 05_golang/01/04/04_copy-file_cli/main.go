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

	f, err := os.Open(toCopy)
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer f.Close()

	f2, err2 := os.Create(newCopy)
	if err2 != nil {
		log.Fatalln("couldn't create file", err2.Error())
	}
	defer f2.Close()

	bs, err3 := ioutil.ReadAll(f)
	if err3 != nil {
		log.Fatalln("couldn't read all", err3.Error())
	}

	_, err4 := f2.Write(bs)
	if err4 != nil {
		log.Fatalln("couldn't write byte slice", err4.Error())
	}

	fmt.Println("Files copied")
}
