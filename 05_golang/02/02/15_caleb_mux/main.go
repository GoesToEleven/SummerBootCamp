package main

import (
	"io"
	"net/http"
	"strings"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "DOG")
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	catName := strings.Split(req.URL.Path, "/")[2]
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    Cat Name: <strong>`+catName+`</strong><br>
    <img src="http://www.vetprofessionals.com/catprofessional/images/home-cat.jpg">
  </body>
</html>`)
}

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cats/", cat)

	http.ListenAndServe(":9000", mux)
}


/*
NOTE:
this file is susceptible to XSS
use escaping of unsafe characters to avoid this
we'll learn more about this later
 */