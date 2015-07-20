package main

import (
//	"encoding/json"
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
//	router.POST("/", updateList)
	http.Handle("/", router)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
}

//func updateList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
//	var item ToDo
//	json.NewDecoder(req.Body).Decode(&item)
//	ctx := appengine.NewContext(req)
//	key := datastore.NewKey(ctx, "ToDo", 0, 0, nil)
//	_, err := datastore.Put(ctx, key, &item)
//	if err != nil {
//		http.Error(res, err.Error(), 500)
//	}
//	// show new list
//	http.Redirect(res, req, "/", 302)
//}

func showList(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("assets/templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	ctx := appengine.NewContext(req)
	qp := datastore.NewQuery("ToDo")
	ToDos := []ToDo{}
	keys, error := qp.GetAll(ctx, &ToDos)
	if error != nil {
		log.Infof(ctx, "%v", error.Error())
	}

	log.Infof(ctx, "HERE IS THE SLICE: %v", ToDos)
	log.Infof(ctx, "HERE IS THE SLICE LENGTH: %v", len(ToDos))
	log.Infof(ctx, "HERE ARE THE KEYS: %v", keys)
	log.Infof(ctx, "HERE IS THE KEYS LENGTH: %v", len(keys))

	ToTres := []ToDo{}
	if (len(ToDos) == 0) {
		ToTres = nil
	} else {
		for i, value := range ToDos {
			task := ToDo{
				Item: string(value),
				Id: string(keys[i]),
			}
			ToTres = append(ToTres, task)
		}
		log.Infof(ctx, "HERE ARE THE ToTres: %v", ToTres)
		log.Infof(ctx, "HERE IS THE ToTres LENGTH: %v", len(ToTres))
//		json.NewEncoder(res).Encode(ToTres)
	}

	err = tpl.ExecuteTemplate(res, "todo-list", ToDos)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}