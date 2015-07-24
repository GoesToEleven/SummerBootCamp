package storage

import (
	"io"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"html/template"
)

const bucketName = "serious-water-88716.appspot.com"
var templates *template.Template

func init() {
	tpl := template.New("tpl")
	templates = template.Must(tpl.ParseGlob("public/*"))

	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/download", handleDownload)
	http.HandleFunc("/list", handleList)
	http.Handle("/", http.FileServer(http.Dir("public/")))
//	http.Handle("/css/", http.FileServer(http.Dir("/")))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))
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

	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(ctx, storage.ScopeFullControl),
			Base:   &urlfetch.Transport{Context: ctx},
		},
	}

	cctx := cloud.NewContext(appengine.AppID(ctx), hc)

	rdr, hdr, err := req.FormFile("form-upload-file")
	if err != nil {
		http.Error(res, "ERROR RECEIVING FILE: "+err.Error(), 500)
		return
	}

	writer := storage.NewWriter(cctx, bucketName, hdr.Filename)
	io.Copy(writer, rdr)

	err = writer.Close()
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

func handleDownload(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	fileName := req.FormValue("fname")

	rdr, err := storage.NewReader(cctx, bucketName, fileName)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	defer rdr.Close()
	res.Header().Set("Content-Disposition", `attachment; filename="` + fileName + `"`)
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

	var fileNames []string
	for _, obj := range objs.Results {
		fileNames = append(fileNames, obj.Name)
	}
	templates.ExecuteTemplate(res, "index", fileNames)
}
