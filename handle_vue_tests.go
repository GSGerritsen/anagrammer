package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func HandleVueTests(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates, err := template.ParseGlob("templates/**/*.html")
	if err != nil {
		panic(err)
	}

	vueTemplate := templates.Lookup("vue-layout.html")
	vueTemplate.Execute(w, nil)
}
