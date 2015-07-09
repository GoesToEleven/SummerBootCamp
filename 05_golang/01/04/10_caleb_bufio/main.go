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


