package GCS

import (
	//	"google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine"
	"golang.org/x/net/context"
)

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	var ctx context.Context = appengine.NewContext(req)
	//	ctx := appengine.NewContext(req)
	session := getSession(ctx, req)

	if req.Method == "POST" {
		session.Bucket = req.FormValue("bucket")
		session.Credentials = req.FormValue("credentials")
		putSession(ctx, res)
		http.Redirect(res, req, "/browse/", 302)
		return
	}

	err := tpls.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func browse(res http.ResponseWriter, req *http.Request) {
	//	ctx := appengine.NewContext(req)
	err := tpls.ExecuteTemplate(res, "browse.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
