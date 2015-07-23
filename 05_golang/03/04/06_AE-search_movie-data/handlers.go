package movies

import (
	"google.golang.org/appengine"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine/search"
	"google.golang.org/appengine/log"
	"encoding/json"
)

// movie struct
type Movies struct {
	Title string
	Description string
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	var err error
	if req.Method == "POST" {
		err = handlePut(ctx, res, req)
	} else if req.Method == "GET" {

	}
	if err != nil {
		http.Error(res, "error", 500)
	}

		res.Header().Set("Content-Type", "text/html")
	renderTemplate(res, "index.html", nil)
}

func handlePut(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
	// post movie to appengine search
	index, err := search.Open("Movies")
	if err != nil {
		return nil, error
	}
	log.Debugf(ctx, "Here is the index: %v", index)
	// instance of movie struct
	var movie Movies
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
	var movies []Movies
	for {
		var movie Movies
		id, err := iterator.Next(&movie)
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