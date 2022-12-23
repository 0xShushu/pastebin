package main

import (
	"html/template"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

type Page struct{
	Code string
}

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

//render to index.html
func MainHandler(w http.ResponseWriter, r *http.Request) {
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

//save code in the server
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	//generate random name
	name := RandString()
	
	//get post parameters
	r.ParseForm()
	code := r.Form.Get("code")

	//write file
	err := os.WriteFile("code/"+name, []byte(code), 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//redirect to the code
	http.Redirect(w, r, "/code/"+name, 301)
}

//display saved code
func CodeHandler(w http.ResponseWriter, r *http.Request) {
	//get code name
	vars := mux.Vars(r)
	name := vars["name"]
	
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