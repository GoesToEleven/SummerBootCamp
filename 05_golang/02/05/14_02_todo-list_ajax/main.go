package main

import (
	"encoding/json"
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
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
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
}

func showList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("assets/templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	ctx := appengine.NewContext(req)
	qp := datastore.NewQuery("ToDo")
	ToDos := []ToDo{}
	_, error := qp.GetAll(ctx, &ToDos)
	if error != nil {
		log.Infof(ctx, "%v", error.Error())
	}

	log.Infof(ctx, "HERE IS THE SLICE: %v", ToDos)
	log.Infof(ctx, "HERE IS THE SLICE LENGTH: %v", len(ToDos))

	if (len(ToDos) == 0) {
		ToDos = nil
	} else {
		json.NewEncoder(res).Encode(ToDos)
	}

	err = tpl.ExecuteTemplate(res, "todo-list", ToDos)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}