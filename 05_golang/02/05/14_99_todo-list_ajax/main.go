package main

import (
	"encoding/json"
	"net/http"
	"google.golang.org/appengine/user"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
	"log"
)

type ToDo struct {
	Item     string
}

func init() {
	router := httprouter.New()
	router.GET("/", showList)
	router.POST("/", updateList)
	http.Handle("/", router)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func updateList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var item ToDo
	json.NewDecoder(req.Body).Decode(&item)
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "ToDo", 0, 0, nil)
	_, err := datastore.Put(ctx, key, &item)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
	// show new list
	http.Redirect(res, req, "/", 302)
}

func showList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("assets/templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	qp := datastore.NewQuery("ToDo")
	ToDos := []ToDo{}
	keys, error := qp.GetAll(ctx, &ToDos)
	for _, value := range ToDos {

	}

	if error != nil {
		log.Infof(ctx, "%v", error.Error())
	}

	json.NewEncoder(res).Encode(ToDos)

	err = tpl.ExecuteTemplate(res, "todo-list", ToDos)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}