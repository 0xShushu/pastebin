package handlers

import (
	"net/http"
	"html/template"
)


//render to 404
func MyNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	
	t, err := template.ParseFiles("templates/404.html")
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