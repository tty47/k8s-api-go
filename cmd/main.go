package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/jrmanes/k8s-api-go/internal/data"
	"github.com/jrmanes/k8s-api-go/internal/server"
	_ "github.com/lib/pq"
)

// Load that function before starts the service
func init() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file... ERROR: ", err)
	}
}

func main() {
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("SERVER_PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt .
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
