package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/memcache"
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", counter)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}


func counter(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	// create the item
	// func Increment(c appengine.Context, key string, delta int64, initialValue uint64) (newValue uint64, err error)
	counter, err := memcache.Increment(ctx, u.Email, 1, 0)
	if err != nil {
		http.Error(res, "counter didn't work", 500)
		return
	}

	fmt.Fprint(res, counter)
}
