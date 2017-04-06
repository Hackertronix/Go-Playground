package main

import (
	"log"
	"os"
	"text/template"
)
8
var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	str := []string{"Steve Jobs", "Steve Wozniak", "Mike Markulla", "Paul Alan"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", str)

	if err != nil {
		log.Fatalln(err)
	}
}
