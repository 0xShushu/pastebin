package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/0xshushu/pastebin/handlers"
)

//server struct
type Server struct {
	Router *chi.Mux
}

//function to initialize a new server
func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

//function to mount handlers
func (s *Server) MountHandlers() {
	//set middleware
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.RedirectSlashes)

	//set routes
	s.Router.Get("/", handlers.IndexHandler)
	s.Router.Post("/save", handlers.SaveHandler)
	s.Router.Get("/code/{name}", handlers.CodeHandler)
	s.Router.NotFound(handlers.MyNotFoundHandler)
}