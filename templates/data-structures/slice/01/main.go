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

	names := []string{"Nigel", "Adrian", "Nathan", "Kristy"}

	err := tpl.Execute(os.Stdout, names)
	if err != nil {
		log.Fatalln(err)
	}
}
