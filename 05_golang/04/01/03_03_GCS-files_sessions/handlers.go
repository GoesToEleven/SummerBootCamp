package GCS

import (
	//	"google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine"
	"golang.org/x/net/context"
	"strings"
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

type browseModel struct {
	Bucket string
	Folder string
}

func browse(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	session := getSession(ctx, req)

	if session.Bucket == "" {
		http.Redirect(res, req, "/", 302)
		return
	}

	folder := strings.SplitN(req.URL.Path, "/", 3)[2]
	model := &browseModel{
		Bucket: session.Bucket,
		Folder: folder,
	}

	err := tpls.ExecuteTemplate(res, "browse.html", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
