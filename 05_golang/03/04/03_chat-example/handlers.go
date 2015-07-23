package chat

import (
	"encoding/json"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/channel"
	"google.golang.org/appengine/user"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// API handles api calls
type API struct {
	root string
}

// NewAPI creates a new API, root should be set to the root url for the API
func NewAPI(root string) *API {
	api := &API{
		root: root,
	}
	return api
}

func (api *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	endpoint := req.URL.Path[len(api.root):]
	method := req.Method

	var err error
	switch endpoint {
	case "channels":
		switch method {
		case "POST":
			err = api.handlePostChannel(res, req)
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	case "messages":
		switch method {
		case "POST":
			err = api.handlePostMessage(res, req)
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	default:
		http.NotFound(res, req)
		return
	}

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}
}

func (api *API) handlePostChannel(res http.ResponseWriter, req *http.Request) error {
	// create channels
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	tok, err := channel.Create(ctx, u.Email)
	if err != nil {
		return err
	}
	err = addConnection(ctx, u.Email)
	if err != nil {
		return err
	}
	// sends token to client
	json.NewEncoder(res).Encode(tok)
	return nil
}

func (api *API) handlePostMessage(res http.ResponseWriter, req *http.Request) error {
	ctx := appengine.NewContext(req)
	// handle incoming messages, send them out to everyone
	type Post struct{
		Text string
	}
	var post Post
//	post := struct{
//	Text string
//	}{}
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		return err
	}
	log.Debugf(ctx, "the post is %s", post.Text)

	query := datastore.NewQuery("connections")
	it := query.Run(ctx)
	for {
		var connection struct{Email string}
		_, err := it.Next(&connection)
		if err == datastore.Done {
			break
		} else if err != nil {
			return err
		}
		log.Debugf(ctx, "here is the email %s", connection.Email)
		err = channel.SendJSON(ctx, connection.Email, post)
		if err != nil {
			return err
		}
	}
	return nil
}

// add new connection
func addConnection(ctx context.Context, email string) error {
	key := datastore.NewKey(ctx, "connections", email, 0, nil)
	type Connections struct {
		Email string
	}
	connection := Connections{
		Email: email,
	}
	_, err := datastore.Put(ctx, key, &connection)
	return err
}

// remove connection
//func removeConnection(ctx context.Context, email string) error {
//	key := datastore.NewKey(ctx, "connections", email, 0, nil)
//	err := datastore.Delete(ctx, key)
//	return err
//}
