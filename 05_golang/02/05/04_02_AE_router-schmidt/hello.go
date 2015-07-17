package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
)

func init() {
	router := httprouter.New()

	router.GET("/:page", handleIndex)

	http.Handle("/", context.ClearHandler(router))
	//http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}

func handleIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// if I go to: localhost:8080/admin
	// page == "admin"
	params.ByName("page")
}
