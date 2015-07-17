package main

import (
	"net/http"
	"github.com/gorilla/context"
	"fmt"
)

func init() {

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", handleIndex)

	http.Handle("/", context.ClearHandler(serveMux))
	//http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, "running serveMux on app engine!")
}
