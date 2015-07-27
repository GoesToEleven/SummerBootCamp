package GCS

import (
	"google.golang.org/appengine"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	//	ctx := appengine.NewContext(req)
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
