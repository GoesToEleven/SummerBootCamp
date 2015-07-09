package main

import (
	"os"
	"log"
	"io/ioutil"
)

func main() {
	// create new file
	f, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalln("couldn't create file!", err.Error())
	}
	defer f.Close()

	writeAt := 1

	for _, v := range os.Args[1:] {
		//open file passed in
		o, err2 := os.Open(v)
		if err2 != nil {
			log.Fatalln("couldn't open a file passed in", err.Error())
		}
		defer o.Close()

		// open returns a *file
		// we need to make it a []byte
		bs, err22 := ioutil.ReadAll(o)
		if err22 != nil {
			log.Fatalln("couldn't make bs", err.Error())
		}

		//add opened file to created file
		n, err3 := f.WriteAt(bs, writeAt)
		if err3 != nil {
			log.Fatalln("couldn't write file", err.Error())
		}
		writeAt += (n+1)
	}
}
