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

func main() {
	//var dog DogHandler
	//var cat CatHandler

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/cats/", func(res http.ResponseWriter, req *http.Request) {
		catName := strings.Split(req.URL.Path, "/")[2]
		io.WriteString(res, `<!DOCTYPE html>
  <html>
    <body>
      Cat Name: <strong>`+catName+`</strong><br>
      <img src="/home-cat.jpg">
    </body>
  </html>`)
	})

	http.ListenAndServe(":9000", nil)
}


