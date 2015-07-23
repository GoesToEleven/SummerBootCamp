package movies

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/search", handleSearch)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}
