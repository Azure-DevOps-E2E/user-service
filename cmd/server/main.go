package main

import (
	"log"
	"os"

	"nexuscart/user-service/internal/server"
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
