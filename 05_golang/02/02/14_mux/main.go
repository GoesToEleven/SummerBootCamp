package main

import (
	"io"
	"net/http"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<img src="http://fillmurray.com/g/200/300" alt="dog">
</body>
</html>`)
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<img src="https://placekitten.com/g/200/300" alt="kitten">
</body>
</html>`)
}

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/dog", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":9000", mux)
}
