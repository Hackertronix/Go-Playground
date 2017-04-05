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

	//template execute requests a writer and os.Stdout is a writer, this can be a file as well
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

}
