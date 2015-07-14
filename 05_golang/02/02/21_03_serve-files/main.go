package main

import (
	"net/http"
	"io"
)


func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.URL.Path)
}

func main() {

	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":9000", nil)
}
