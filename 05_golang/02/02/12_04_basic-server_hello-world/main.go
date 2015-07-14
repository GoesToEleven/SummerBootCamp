package main

import (
	"io"
	"net/http"
	"strings"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if strings.Contains(req.URL.Path, "Cat") {
		io.WriteString(res, "CAT")
	} else {
		io.WriteString(res, "DOG")
	}

	io.WriteString(res, req.RequestURI)
}

func main() {

	var h myHandler
	http.ListenAndServe(":9000", h)
}
