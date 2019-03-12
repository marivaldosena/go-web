package main

import (
	"log"
	"os"
	"path"
	"text/template"
)

var tpl *template.Template
var arquivo string

func init() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo = path.Join(caminho, "templates", "index.html")

	tpl = template.Must(template.ParseFiles(arquivo))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", sages)

	if err != nil {
		log.Fatalln(err)
	}
}
