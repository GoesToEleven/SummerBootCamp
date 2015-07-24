package storage

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

const bucketName = "golang-bootcamp"

func init() {
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/list", handleList)
}

func getCloudContext(appengineContext context.Context) context.Context {
	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(appengineContext, storage.ScopeFullControl),
			Base:   &urlfetch.Transport{Context: appengineContext},
		},
	}

	return cloud.NewContext(appengine.AppID(appengineContext), hc)
}

func handlePut(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	writer := storage.NewWriter(cctx, bucketName, "example.txt")
	io.WriteString(writer, "Works")
	err := writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	rdr, err := storage.NewReader(cctx, bucketName, "example.txt")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	defer rdr.Close()

	io.Copy(res, rdr)
}

func handleList(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	var query *storage.Query
	objs, err := storage.ListObjects(cctx, bucketName, query)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	for _, obj := range objs.Results {
		fmt.Fprintln(res, obj.Name)
	}
}
