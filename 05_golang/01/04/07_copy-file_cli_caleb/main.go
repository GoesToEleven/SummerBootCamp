package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func cp(srcFileName, dstFileName string) error {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("error writing destination file: %v", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: my-cp <SRC> <DST>")
	}
	srcFileName := os.Args[1]
	dstFileName := os.Args[2]

	err := cp(srcFileName, dstFileName)
	if err != nil {
		log.Fatalln(err)
	}
}
