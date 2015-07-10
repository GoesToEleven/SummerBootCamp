package main

import (
	"os"
)

func main() {
	src, err := os.Open("src-file")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst1, err := os.Create("dst1-file")
	if err != nil {
		panic(err)
	}
	defer dst1.Close()

	bs := make([]byte, 5)
	src.Read(bs)
	dst1.Write(bs)

}