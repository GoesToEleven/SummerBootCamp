package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	bufferedReader := bufio.NewReader(srcFile)
	io.Copy(someDestination, bufferedReader)
}

// might not be a complete working file
// this file was just pulled part way through lecture
// https://www.youtube.com/playlist?list=PLSak_q1UXfPqJ6pIpJt2oNJTPx9ByBy6K


