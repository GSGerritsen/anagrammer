package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RenderTemplate(w, r, "index/home", nil)
}

func HandleWordSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	word := r.FormValue("word")

	anagrams, err := SearchDBForAnagrams(word)

	if err != nil {
		// Render an error in the template instead
		panic(err)
	}

	RenderTemplate(w, r, "index/home", map[string]interface{}{
		"WordsEnglish": anagrams,
	})
}
