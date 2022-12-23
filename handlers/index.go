package handlers

import (
	"html/template"
	"net/http"
)


//render to index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err2 := t.Execute(w, nil)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}