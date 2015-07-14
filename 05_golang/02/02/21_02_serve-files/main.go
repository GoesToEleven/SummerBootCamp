package main

import (
	"net/http"
	"io"
)

type myDog int

func (h myDog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.URL.Path)
}

func main() {
	var dog myDog

	http.Handle("/dog/", dog)
	http.ListenAndServe(":9000", nil)
}
