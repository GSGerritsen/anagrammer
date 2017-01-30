package main

import (
	//"encoding/json"
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

func HandleHomeGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	anagramsListJSON, err := SearchDBForAnagrams(params.ByName("word"))
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(anagramsListJSON)
}
