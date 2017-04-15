package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	str := map[string]string{

		"Machine": "Mac",
		"Tune":    "Linkin Park",
		"Food":    "Continental",
		"Love":    "Work"}

	err := tpl.Execute(os.Stdout, str)
	if err != nil {
		log.Fatalln(err)
	}
}
