package api

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
}

// NewServer ...
func NewServer() *Server {
	return &Server{}
}

// Start ...
func (s *Server) Start() {
	router := mux.NewRouter()

	// Add all routes to the rotuer
	for _, route := range Routes {
		router.NewRoute().Name(route.Name).Methods(route.Method).Path(route.Pattern).HandlerFunc(route.HandlerFunc)
	}

	// Add all middlewares
	neg := negroni.Classic()
	for _, middleware := range Middlewares {
		neg.UseFunc(middleware)
	}

	neg.UseHandler(router)

	address := "0.0.0.0:8080"
	log.Println("Starting server at ", address)
	log.Fatal(http.ListenAndServe(address, neg))
}
