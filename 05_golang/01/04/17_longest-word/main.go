package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func LongestWord(rdr io.Reader) (string, int) {

	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	var longestWord string
	longestWordLength := 0
	for scanner.Scan() {
		word := scanner.Text()
		// insert code here
		if len(word) > longestWordLength {
			longestWordLength = len(word)
			longestWord = word
		}
	}
	// return something
	return longestWord, longestWordLength
}

func main() {
	srcFile, err := os.Open("../resources/moby.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	longestWord, wordsLength := LongestWord(srcFile)
	fmt.Println("longest word:", longestWord)
	fmt.Println("the words length:", wordsLength)
}
