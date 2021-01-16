package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
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
	fmt.Println("main")
}
