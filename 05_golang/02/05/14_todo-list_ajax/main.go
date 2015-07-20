package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/user"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
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
	key := datastore.NewKey(ctx, "ToDo", u.Email, 0, nil)
	_, err := datastore.Put(ctx, key, &item)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func showList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("assets/templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(res, "todo-form", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}


func getAPIProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	var profile Profile
	err := datastore.Get(ctx, key, &profile)
	if err != nil {
		profile.Email = u.Email
	}

	json.NewEncoder(res).Encode(profile)
}


func showIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "/profile", 302)
}

