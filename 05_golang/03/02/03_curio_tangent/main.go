package main

import (
	"net/http"
	"io"
)

type ArticleHandler int

func (a ArticleHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is the article")
}

type UsersHandler int

func (h UsersHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is info about users")
}

func main() {

	var user UsersHandler
	var article ArticleHandler


	mux := http.NewServeMux()

	mux.Handle("/articles/", article)

	mux.Handle("/users", user)

	http.ListenAndServe(":3000", mux)

}