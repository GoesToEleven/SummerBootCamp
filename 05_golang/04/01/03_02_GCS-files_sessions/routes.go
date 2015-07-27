package GCS

import (
	"html/template"
	"net/http"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", index)
	http.HandleFunc("/browse/", browse)
}
