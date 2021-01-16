package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	v1 "github.com/jrmanes/k8s-api-go/internal/server/v1"
)

type Server struct {
	server *http.Server
}

// New initialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	r.Mount("/", v1.New())

	// validate the port received, if not, use default 8080
	if len(port) == 0 {
		port = "8080"
	}

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
