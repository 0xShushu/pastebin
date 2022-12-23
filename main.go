package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"os"
)

func main() {
	//making a new router
	r := mux.NewRouter().StrictSlash(true)
	//set routes
	r.HandleFunc("/", MainHandler).Methods("GET")
	r.HandleFunc("/save", SaveHandler).Methods("POST")
	r.HandleFunc("/code/{name}", CodeHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(MyNotFoundHandler)

	//get port for http server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//listen and serve
	fmt.Println(http.ListenAndServe(":"+port, r))
}