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

	rdr := io.TeeReader(src, dst1)
	rdr = io.TeeReader(rdr, os.Stdout)

	io.Copy(dst2, rdr)

}

/*
WHAT TEEREADER DOES....

func TeeReader(r Reader, w Writer) Reader

func teeReader(rdr io.Reader) {
	bs := make([]byte, 10)
	rdr.Read(bs)
	w.Write(bs)
	return bs


}

 */