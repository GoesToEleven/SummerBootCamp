package main

import (
	"io"
	"net/http"
	"os"
)

func uploadPage(res http.ResponseWriter, req *http.Request) {
	// CHECK FOR VALID SESSION
	// you need to put the "check for valid session code" in here

	if req.Method == "POST" {
		// <input type="file" name="file">
		src, hdr, err := req.FormFile("file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		// create a new file
		// make sure you have a "tmp" directory in your web root
		dst, err := os.Create("tmp/" + hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		// copy the uploaded file into the new file
		io.Copy(dst, src)
	}

	// RENDER THE UPLOAD PAGE
}
