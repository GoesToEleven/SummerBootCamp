package GCS

import (
	//	"google.golang.org/appengine"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"google.golang.org/appengine/memcache"
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

	if req.Method == "POST" {
		// set cookie
		sessionID, _ := uuid.NewV4()

		cookie := &http.Cookie{
			Name:  "Session",
			Value: sessionID.String(),
		}
		http.SetCookie(res, cookie)
		// set memcache
		item1 := &memcache.Item{
			Key:   sessionID.String(),
			Value: []byte("bar"), // TODO: Get Value From FORM
		}
		if err := memcache.Set(c, item1); err != nil {
			return err
		}
	}
}

func browse(res http.ResponseWriter, req *http.Request) {
	//	ctx := appengine.NewContext(req)
	err := tpls.ExecuteTemplate(res, "browse.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
