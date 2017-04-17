package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Name  string
	Model string
	Doors int
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	b := sage{"Buddha", "The belief is no belief"}
	g := sage{"Gandhi", "Be change"}

	f := car{
		Name:  "Ford",
		Model: "F150",
		Doors: 2,
	}

	c := car{
		Name:  "Toyota",
		Model: "Corolla",
		Doors: 4,
	}

	sages := []sage{b, g}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}

}
