package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func totalWords(rdr io.Reader) int {

	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	totalWords := 0
	for scanner.Scan() {
		totalWords++
	}
	// return something
	return totalWords
}

func main() {
	srcFile, err := os.Open("../resources/moby.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	totalWords := totalWords(srcFile)
	fmt.Println("total words: ", totalWords)
	fmt.Println("total pages: ", totalWords/250)
}
