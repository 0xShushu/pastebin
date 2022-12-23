package handlers

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"os"
	"html/template"
)


//display saved code
func CodeHandler(w http.ResponseWriter, r *http.Request) {
	//get code name
	name := chi.URLParam(r, "name")
	
	//read code
	codeByte, err := os.ReadFile("code/"+name)	
	if err != nil {
		http.Error(w, "ID not found", 404)
		return
	}
	
	//display code
	page := Page{string(codeByte)}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, page)
}