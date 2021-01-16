package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	v1 "github.com/jrmanes/k8s-api-go/internal/server/v1"
)

// Server define an http standard library
type Server struct {
	server *http.Server
}

// New initialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	// use chi middleware in order to log request request to our API
	r.Use(middleware.Logger)

	r.Mount("/", v1.New())

	// validate the port received, if not, use default 8080
	if len(port) == 0 {
		port = "8080"
	}

	// set configuration to our server
	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
