package GCS
import (
	"golang.org/x/net/context"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/appengine"
	"google.golang.org/cloud/storage"
	"io/ioutil"
)

func getCloudContext(aeCtx context.Context) (context.Context, error) {
	data, err := ioutil.ReadFile("/Users/user_name/Desktop/json_file.json")
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON (
		data,
		storage.ScopeFullControl,
	)

	tokenSource := conf.TokenSource(aeCtx)

	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
			Base:   &urlfetch.Transport{Context: aeCtx},
		},
	}

	return cloud.NewContext(appengine.AppID(aeCtx), hc), nil
}
