package main

import (
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

type ToDo struct {
	Item     string
	Id 		 string
}

func init() {
	router := httprouter.New()
	router.GET("/", showList)
	http.Handle("/", router)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func showList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("assets/templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(res, "todo-list", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}