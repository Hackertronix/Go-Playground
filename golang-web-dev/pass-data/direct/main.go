package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)

	if err != nil {
		log.Fatal(err)
	}

}
