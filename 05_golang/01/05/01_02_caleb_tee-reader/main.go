package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("src-file")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst1, err := os.Create("dst1-file")
	if err != nil {
		panic(err)
	}
	defer dst1.Close()

	dst2, err := os.Create("dst2-file")
	if err != nil {
		panic(err)
	}
	defer dst2.Close()

	dst3, err := os.Create("dst3-file")
	if err != nil {
		panic(err)
	}
	defer dst3.Close()

	rdr := io.TeeReader(src, dst1)
	rdr = io.TeeReader(rdr, os.Stdout)
	rdr = io.TeeReader(rdr, dst2)

	io.Copy(dst3, rdr)

}

/*
WHAT TEEREADER DOES....

func TeeReader(r Reader, w Writer) Reader

TeeReader returns a Reader that writes to w what it reads from r. All reads from r performed through it are matched with corresponding writes to w. There is no internal buffering - the write must complete before the read completes. Any error encountered while writing is reported as a read error.


 */