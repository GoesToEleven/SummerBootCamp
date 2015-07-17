package main

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"fmt"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", profile)
	http.HandleFunc("/show/", showUser)
}

type User struct {
	FirstName       string
	LastName string
}


func profile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, `
		<form action="/" method="post">
			<input type="text" id="fname" name="fname">
			<input type="text" id="fname" name="lname">
			<input type="submit">
		</form>
	`)

	if req.Method == "POST" {
		email := u.Email
		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		fmt.Println(email, fname, lname)
		key := datastore.NewKey(ctx, "User", email, 0, nil)
		entity := &User{
			FirstName:       fname,
			LastName: lname,
		}

		_, err := datastore.Put(ctx, key, entity)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		http.Redirect(res, req, "/", 302)

	}

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
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<dl>
			<dt>`+entity.FirstName+`</dt>
			<dd>`+entity.LastName+`</dd>
		</dl>
	`)
}
