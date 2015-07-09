package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func LongestWord(rdr io.Reader) string {

	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		// insert code here
	}
	// return something
}

func main() {
	srcFile, err := os.Open("moby.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	counts := LongestWord(srcFile)
	fmt.Println("Number of whales:", counts["whale"])
}
