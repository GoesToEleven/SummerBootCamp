package movies

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
	"golang.org/x/net/context"
	"github.com/goestoeleven/csuf/sp15/xx_in_process/06_metacasts/not_gone_through/go_appengine/goroot/src/appengine/search"
	"google.golang.org/appengine/log"
)

func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
//	u := user.Current(ctx)

	if req.Method == "POST" {
		err := handlePut(ctx, res, req)
	}

	if req.Method == "GET" {

	}

		res.Header().Set("Content-Type", "text/html")
	 render the login template
	renderTemplate(res, "index.html", nil)
}

func handlePut(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
	// post movie to appengine search
	index, err := search.Open("Movies")
	if err != nil {
		return nil, error
	}
	log.Debugf(ctx, "Here is the index: %v", index)
	// movie struct
	type Movies struct {
		Title string
		Description string
	}
	// instance of movie struct
	var movie Movies
	// receive JSON from form

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

func handleGet(res http.ResponseWriter, req *http.Request) {

}

func handleSearch(res http.ResponseWriter, req *http.Request) {

}
