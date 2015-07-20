package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
	"html/template"
	consoleLog "log"
	"net/http"
	"reflect"
)

var tpl *template.Template

type Profile struct {
	Email    string
	UserName string
}

func init() {
	var err error
	tpl, err = template.ParseFiles("templates/html/header.html",
		"templates/html/footer.html",
		"templates/html/home.html",
		"templates/html/login.html",
		"templates/html/profile.html")
	if err != nil {
		consoleLog.Fatalln("couldn't parse", err.Error())
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/profile", profile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
}

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	err := tpl.ExecuteTemplate(res, "home.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}

}

func login(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	err := tpl.ExecuteTemplate(res, "login.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}

	log.Infof(ctx, "HERE IS THE InfoF USER: %v", u)
	log.Errorf(ctx, "HERE IS ErrorrF THE USER: %v", u)

	// see if email has associated username
	// ask for username if doesn't exist
	// make sure username is unique
	q := datastore.NewQuery("Profiles").Filter("Email =", u.Email)
	iterator := q.Run(ctx)
	for {
		var user Profile
		key, err := iterator.Next(&user)
		if err == datastore.Done {
			// zero records returned
			// no profile in datastore
			// redirect to create profile in datastore
			http.Redirect(res, req, "/profile", 301)
			break
		} else if err != nil {
			log.Errorf(ctx, "error retrieving user profile: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		//log.Infof(ctx, "USER ID: %v", user.ID)
		log.Infof(ctx, "USER EMAIL: %v", user.Email)
		log.Infof(ctx, "USER USERNAME: %v", user.UserName)
		log.Infof(ctx, "RETURNED KEY: %v", key.IntID())
		//user.ID = key.IntID()
	}
}

func profile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
//	u := user.Current(ctx)

	// ask for username
	// make sure username is unique
	log.Infof(ctx, "USER NAME: %v", req.FormValue("user-name") )
	log.Infof(ctx, "USER NAME TYPEOF: %v", reflect.TypeOf(req.FormValue("user-name")) )
	if req.FormValue("user-name") != "" {
		q := datastore.NewQuery("Profiles").Filter("UserName =", req.FormValue("user-name"))
		//q.Count()

		iterator := q.Run(ctx)
		for {
			var user Profile
			key, err := iterator.Next(&user)
			if err == datastore.Done {
				// zero records returned
				// no profile in datastore
				// redirect to create profile in datastore
				http.Redirect(res, req, "/profile", 302)
				break
			} else if err != nil {
				log.Errorf(ctx, "error retrieving user profile: %v", err)
				http.Error(res, err.Error(), 500)
				return
			}
			//log.Infof(ctx, "USER ID: %v", user.ID)
			log.Infof(ctx, "USER EMAIL: %v", user.Email)
			log.Infof(ctx, "USER USERNAME: %v", user.UserName)
			log.Infof(ctx, "RETURNED KEY: %v", key.IntID())
			//user.ID = key.IntID()
		}
	}

	err := tpl.ExecuteTemplate(res, "profile.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}

}
//
//func getProfileByEmail(res http.ResponseWriter, req *http.Request) (*Profile, error) {
//	ctx := appengine.NewContext(req)
//	u := user.Current(ctx)
//	q := datastore.NewQuery("Profiles").Filter("Email =", u.Email)
//	q.Count()
//}
//
//func getProfileByUsername(res http.ResponseWriter, req *http.Request) (*Profile, error) {
//	ctx := appengine.NewContext(req)
//	u := user.Current(ctx)
//}
