package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	log.SetFlags(0)

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, []int{1,2,3,4,5,})
	if err != nil {
		log.Fatalln(err)
	}
}


