package main

import (
	"net/http"
	"fmt"
	"os"
	"github.com/0xshushu/pastebin/server"
)

func main() {
	//making a new server
	s := server.NewServer()
	
	//mount handlers
	s.MountHandlers()

	//get port for http server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//listen and serve
	fmt.Println(http.ListenAndServe(":"+port, s.Router))
}