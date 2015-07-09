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


// might not be a complete working file
// this file was just pulled part way through lecture
// https://www.youtube.com/playlist?list=PLSak_q1UXfPqJ6pIpJt2oNJTPx9ByBy6K