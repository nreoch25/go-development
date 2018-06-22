package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

func init() {
	tpl = template.Must(tpl.ParseFiles("index.gohtml"))
}
