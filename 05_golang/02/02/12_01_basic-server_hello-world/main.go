package main

import (
	"net/http"
	"fmt"
)

type myHandler int

func (h myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Hello world")
}

func main() {
	var h myHandler
	http.ListenAndServe(":9000", h)
}