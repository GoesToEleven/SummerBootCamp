package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
	})

	http.HandleFunc("/home-cat.jpg", func(res http.ResponseWriter, req *http.Request) {
		f, err := os.Open("home-cat.jpg")
		if err != nil {
			http.Error(res, "sorry no cat", 404)
			return
		}
		defer f.Close()

		io.Copy(res, f)

	})

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


