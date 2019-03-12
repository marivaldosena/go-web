package main

import (
	"log"
	"os"
	"path"
	"text/template"
)

var tpl *template.Template

func init() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo := path.Join(caminho, "templates", "index.html")

	tpl = template.Must(template.ParseFiles(arquivo))
}

func main() {
	paises := map[string]string{
		"BRA": "Brazil",
		"PRT": "Portugal",
		"AGO": "Angola",
		"MOZ": "Mozambique",
		"ARG": "Argentina",
		"URY": "Uruguay",
		"PRY": "Paraguay",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", paises)

	if err != nil {
		log.Fatalln(err)
	}
}
