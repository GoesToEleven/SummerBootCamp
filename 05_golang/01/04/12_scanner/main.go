package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
//		fmt.Println(">>>", line)
//		fmt.Println(strings.ToUpper(line[0:1]) + strings.ToLower(line[1:]))
		fmt.Println(strings.ToUpper(string(line[0])) + strings.ToLower(line[1:]))
	}
}