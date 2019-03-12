package main

import (
	"log"
	"os"
	"path"
	"text/template"
)

var tpl *template.Template
var arquivo string
var err error

func init() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo = path.Join(caminho, "templates", "index.html")
	tpl = template.Must(template.ParseFiles(arquivo))
}

func main() {
	err = tpl.ExecuteTemplate(os.Stdout, "index.html", "Teste")
	if err != nil {
		log.Fatalln(err)
	}
}
