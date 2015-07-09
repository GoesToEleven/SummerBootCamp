package main

import (
	"log"
	"os"
	"strings"
	"io"
)

func main() {
	dstFile, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file: %v", err)
	}
	defer dstFile.Close()

	io.Copy(dstFile, strings.NewReader("hello world"))
	dstFile.Write([]byte("hello world"))
}