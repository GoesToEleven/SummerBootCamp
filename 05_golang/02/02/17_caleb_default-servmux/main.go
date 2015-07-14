package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "DOG")
	})
	http.HandleFunc("/cats/", func(res http.ResponseWriter, req *http.Request) {
		catName := strings.Split(req.URL.Path, "/")[2]
		io.WriteString(res, `<!DOCTYPE html>
  <html>
    <body>
      Cat Name: <strong>`+catName+`</strong><br>
      <img src="http://www.vetprofessionals.com/catprofessional/images/home-cat.jpg">
    </body>
  </html>`)
	})

	http.ListenAndServe(":9000", nil)
}