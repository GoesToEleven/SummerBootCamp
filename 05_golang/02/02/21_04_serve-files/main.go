package main

import (
	"net/http"
)


func img(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../../../../images/01.jpeg")
}

func main() {

	http.HandleFunc("/img/", img)
	http.ListenAndServe(":9000", nil)
}
