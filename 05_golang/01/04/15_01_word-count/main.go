package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WordCount(book string) {

	srcFile, err := os.Open(book)
	if err != nil {
		log.Fatalln("error", err, err.Error())
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	searchIn := "../resources/moby.txt"
	WordCount(searchIn)
}

/*
srcFile, err := os.Open(book)
returns *File

func (*File) Read --> means that *File implements Reader interface

type Reader interface {
	Read(p []byte) (n int, err error)
}

bufio.NewScanner
takes a Reader interface::

func NewScanner(r io.Reader) *Scanner

*/

/*
Remember WordCount?
Use WordCount to count the frequency of words for a large book
(eg http://www.gutenberg.org/files/2701/old/moby10b.txt).
*/
