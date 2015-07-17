package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
	"net/http"
)

func init() {
	http.HandleFunc("/", profile)
	http.HandleFunc("/show/", showUser)
}

type User struct {
	FirstName string
	LastName  string
}

func profile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	if req.Method == "POST" {
		email := u.Email
		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		log.Infof(ctx, "RETRIEVED: %v, %v", fname, lname)

		key := datastore.NewKey(ctx, "User", email, 0, nil)
		entity := &User{
			FirstName: fname,
			LastName:  lname,
		}

		log.Infof(ctx, "PUT: %v", entity)
		_, err := datastore.Put(ctx, key, entity)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		http.Redirect(res, req, "/show/", 302)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, `
		<form action="/" method="POST">
			<input type="text" id="fname" name="fname">
			<input type="text" id="lname" name="lname">
			<input type="submit">
		</form>
	`)

}

func showUser(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	email := u.Email

	key := datastore.NewKey(ctx, "User", email, 0, nil)
	var entity User
	err := datastore.Get(ctx, key, &entity)

	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	log.Infof(ctx, "%v", entity)
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<dl>
			<dt>`+entity.FirstName+`</dt>
			<dd>`+entity.LastName+`</dd>
			<dd>`+u.Email+`</dd>
		</dl>
	`)
}
