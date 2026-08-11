package main

import (
	"log"
	"os"

	"polyglot-shop/user-service/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := server.New().Run(":" + port); err != nil {
		log.Fatalf("user-service stopped: %v", err)
	}
}
