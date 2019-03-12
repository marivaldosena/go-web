package main

import (
	"log"
	"os"
	"path"
	"text/template"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	tpl, err := template.ParseFiles(
		path.Join(cwd, "templates", "index.html"))

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
