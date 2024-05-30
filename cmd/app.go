package main

import (
	"go-service-template/internal/router"
	"log"
)

func main() {
	if err := router.Serve().Listen(":3071"); err != nil {
		log.Fatalf("Failed to listen ")
	}
}
