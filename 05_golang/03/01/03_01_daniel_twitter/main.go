package twitter

import (
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type profile struct {
	Username string
}

type tweet struct {
	Message    []string
	SubmitTime time.Time
}

type mainpageData struct {
	Tweets   []tweet
	Logged   bool
	Email    string
	LoginURL string
}

var tpl = template.New("templates")

func init() {
	_, err := tpl.ParseFiles("templates/index.gohtml", "templates/createProfile.gohtml")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handle)
	http.HandleFunc("/CreateProfile", createProfile)
}

func createProfile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	if req.Method == "POST" {
		username := req.FormValue("username")
		// TODO Confirm input is valid
		// TODO Make sure username is not taken
		key := datastore.NewKey(ctx, "profile", u.Email, 0, nil)
		p := profile{
			Username: username,
		}
		_, err := datastore.Put(ctx, key, &p)
		if err != nil {
			http.Error(res, "Server error!", http.StatusInternalServerError)
			log.Errorf(ctx, "Create profile Error: %s\n", err.Error())
			return
		}
	}
	err := tpl.ExecuteTemplate(res, "createProfile.gohtml", nil)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

func handle(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	if u != nil {
		key := datastore.NewKey(ctx, "profile", u.Email, 0, nil)
		var p profile
		err := datastore.Get(ctx, key, &p)
		if err == datastore.ErrNoSuchEntity {
			http.Redirect(res, req, "/CreateProfile", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(res, "Server error!", http.StatusInternalServerError)
			log.Errorf(ctx, "Datastore get Error: %s\n", err.Error())
			return
		}
	}

	// Get recent tweets
	query := datastore.NewQuery("Tweets")
	tweets := []tweet{}
	_, err := query.GetAll(ctx, &tweets)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Query Error: %s\n", err.Error())
		return
	}

	// Create template
	loginURL, err := user.LoginURL(ctx, "/")
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Login URL Error: %s\n", err.Error())
		return
	}
	data := mainpageData{
		Tweets:   tweets,
		Logged:   u != nil,
		LoginURL: loginURL,
	}
	if data.Logged {
		data.Email = u.Email
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}
