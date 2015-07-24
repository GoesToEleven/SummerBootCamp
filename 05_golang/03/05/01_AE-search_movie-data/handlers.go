package movies

import (
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
	"net/http"
)

// movie struct
type Movie struct {
	Title       string
	Description string
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "index.html", nil)
}

func handleApi(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	var err error
	if req.Method == "POST" {
		err = handlePut(ctx, res, req)
	} else if req.Method == "GET" {
		err = handleGet(ctx, res, "")
	}
	if err != nil {
		http.Error(res, "error", 500)
	}
}

func handlePut(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
	// post movie to appengine search
	index, err := search.Open("Movies")
	if err != nil {
		return nil, error
	}
	// log.Debugf(ctx, "Here is the index: %v", index)
	// instance of movie struct
	var movie Movie
	// receive JSON from form
	err = json.NewDecoder(req.Body).Decode(&movie)
	if err != nil {
		return err
	}
	// put the movie into appengine search
	index.Put(ctx, "", &movie)
	// return the new list of movies to client
	err = handleGet(ctx, res, "")
	if err != nil {
		return err
	}
	return nil
}

func handleGet(ctx context.Context, res http.ResponseWriter, criteria string) error {
	index, err := search.Open("Movies")
	if err != nil {
		return err
	}
	var iterator *search.Iterator
	if criteria == "" {
		iterator = index.List(ctx, nil)
	} else {
		iterator = index.Search(ctx, criteria, nil)
	}
	// slice of movie instances
	var movies []Movie
	for {
		var movie Movie
		_, err := iterator.Next(&movie)
		if err == search.Done {
			break
		}
		if err != nil {
			return err
		}
		movies = append(movies, movie)
	}
	json.NewEncoder(res).Encode(movies)
}
