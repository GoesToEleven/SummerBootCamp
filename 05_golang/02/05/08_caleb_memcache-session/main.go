package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"

	"github.com/nu7hatch/gouuid"
)

func init() {

	http.HandleFunc("/", handleIndex)

}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("sessionid")
	if cookie == nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: id.String(),
		}
	}
	http.SetCookie(res, cookie)

	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, cookie.Value)
	if item == nil {
		m := map[string]string{
			"email": "test@example.com",
		}
		bs, _ := json.Marshal(m)

		item = &memcache.Item{
			Key:   cookie.Value,
			Value: bs,
		}
		memcache.Set(ctx, item)
	}

	var m map[string]string
	json.Unmarshal(item.Value, &m)

	fmt.Fprintln(res, m)

}
