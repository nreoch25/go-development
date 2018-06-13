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

	names := map[string]string{"Name1": "Nigel", "Name2": "Adrian", "Name3": "Nathan", "Name4": "Kristy"}

	err := tpl.Execute(os.Stdout, names)
	if err != nil {
		log.Fatalln(err)
	}
}
