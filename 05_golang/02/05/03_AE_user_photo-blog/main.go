package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

var tpl *template.Template


func init() {
	var err error
	tpl, err = template.ParseFiles("assets/tpl/index.gohtml", "assets/tpl/admin.gohtml", "assets/tpl/profile.gohtml", "assets/tpl/admin_upload.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse", err)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/admin/upload", upload)
	http.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs/", http.FileServer(http.Dir("assets/imgs/"))))
}

func home(res http.ResponseWriter, req *http.Request) {

	type Photo struct {
		PhotoPath string
		Lat float64
		Long float64
	}

	var model struct {
		Photos []Photo
		LoggedIn bool
	}

	filepath.Walk("assets/imgs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			var currentPhoto Photo
			currentPhoto.PhotoPath = path

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			exif.RegisterParsers(mknote.All...)
			x, err := exif.Decode(f)
			if err != nil {
				log.Println("no info")
				return nil
			}
			currentPhoto.Lat , currentPhoto.Long, _ = x.LatLong()
			fmt.Println("lat, long: ", currentPhoto.Lat , ", ", currentPhoto.Long)

			model.Photos = append(model.Photos, currentPhoto)
		}
		return nil
	})

	err := tpl.ExecuteTemplate(res, "index.gohtml", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func admin(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	tpl.ExecuteTemplate(res, "admin.gohtml", u.Email)
}


func profile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	tpl.ExecuteTemplate(res, "profile.gohtml", u.Email)
}

func upload(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	// if they are uploading a file, handle that
	if req.Method == "POST" {
		src, hdr, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer src.Close()

		fileName := hdr.Filename
		dst, err := os.Create("assets/imgs/" + fileName)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
	}

	// execute template
	tpl.ExecuteTemplate(res, "admin_upload.gohtml", u.Email)
}
