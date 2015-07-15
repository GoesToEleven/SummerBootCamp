package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

func main() {
	tpls := template.New("templates")

	tpls, err := tpls.ParseFiles("login.gohtml", "logout.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse templates", err, err.Error())
	}

	http.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
		// set cookie
		cookie, err := req.Cookie("my-cookie")
		if err == http.ErrNoCookie {
			cookie = &http.Cookie {
				Name: "my-cookie",
				Value: "Logged In",
			}
		}
		http.SetCookie(res, cookie)
		// execute template
		err = tpls.ExecuteTemplate(res, "login.gohtml", cookie.Value)
		if err != nil {
			log.Fatalln("couldn't respond", err, err.Error())
		}
	})

	http.HandleFunc("/logout/", func(res http.ResponseWriter, req *http.Request) {
		logout := req.FormValue("logout")
		fmt.Println(logout)
		if logout == "Submit" {
			cookie, err := req.Cookie("my-cookie")
			if err == http.ErrNoCookie {
				cookie = &http.Cookie {
					Name: "my-cookie",
					Value: "Logged Out",
				}
			}
		}

		err := tpls.ExecuteTemplate(res, "logout.gohtml", nil)
		if err != nil {
			log.Fatalln("couldn't respond", err, err.Error())
		}
	})

	http.ListenAndServe(":9000", nil)
}
