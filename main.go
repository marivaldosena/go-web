package main

import (
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

var tpl *template.Template

type country struct {
	Name string
}

func toUpper(country string) string {
	return country[:3]
}

var fn = template.FuncMap{
	"uc":    strings.ToUpper,
	"slice": toUpper,
}

func init() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo := path.Join(caminho, "templates", "index.html")

	tpl = template.Must(template.New("").Funcs(fn).ParseFiles(arquivo))
}

func main() {
	countries := []country{
		country{"Brazil"},
		country{"Portugal"},
		country{"Argentina"},
		country{"Uruguay"},
		country{"Paraguay"},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", countries)

	if err != nil {
		log.Fatalln(err)
	}
}
