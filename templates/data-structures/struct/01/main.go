package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type people struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	nigel := people{
		Name:  "Nigel",
		Motto: "Get it done",
	}

	err := tpl.Execute(os.Stdout, nigel)
	if err != nil {
		log.Fatalln(err)
	}
}
