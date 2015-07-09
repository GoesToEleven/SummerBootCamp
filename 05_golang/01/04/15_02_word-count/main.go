package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func WordCount(book string) {

	srcFile, err := os.Open(book)
	if err != nil {
		log.Fatalln("error", err, err.Error())
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	searchIn := "../resources/moby.txt"
	WordCount(searchIn)
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
