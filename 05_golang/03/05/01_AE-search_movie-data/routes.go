package movies

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api", handleApi)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
}
