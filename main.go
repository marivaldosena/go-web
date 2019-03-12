package main

import (
	"html/template"
	"log"
	"os"
	"path"
)

var tpl *template.Template

type country struct {
	Name  string
	Motto string
}

func init() {
	caminho, err := os.Getwd()

	verifyError(err)

	arquivo := path.Join(caminho, "templates", "index.html")
	tpl = template.Must(template.ParseFiles(arquivo))
}

func main() {
	country := country{
		Name:  "Brasil",
		Motto: "Ordem e Progresso",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", country)

	verifyError(err)
}

func verifyError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
