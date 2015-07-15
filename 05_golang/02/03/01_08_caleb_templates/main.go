package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	log.SetFlags(0)

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, Page{
		Title: "My Title 2",
	})
	if err != nil {
		log.Fatalln(err)
	}
}



