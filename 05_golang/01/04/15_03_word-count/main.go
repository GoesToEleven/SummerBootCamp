package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func WordCount(searchIn, searchFor string) int {

	srcFile, err := os.Open(searchIn)
	if err != nil {
		log.Fatalln("error", err, err.Error())
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)

	scanner.Split(bufio.ScanWords)

	var result int
	for scanner.Scan() {
		if searchFor == strings.ToLower(scanner.Text()) {
			result++
		}
	}
	return result
}

func main() {
	searchIn := "moby10b.txt"
	searchFor := "and"
	fmt.Println(WordCount(searchIn, searchFor))
}

/*
see here:
http://godoc.org/bufio#Scanner
*/

/*
Remember WordCount?
Use WordCount to count the frequency of words for a large book
(eg http://www.gutenberg.org/files/2701/old/moby10b.txt).
*/
