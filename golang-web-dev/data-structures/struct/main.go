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

type sage struct {
	Name  string
	Motto string
}

func main() {
	jobs := sage{"Steve", "Stay Hungry Stay Foolish"}

	err := tpl.Execute(os.Stdout, jobs)
	if err != nil {
		log.Fatalln(err)
	}
}
