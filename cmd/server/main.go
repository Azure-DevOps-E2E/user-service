package main

import (
	"log"
	"os"

	"nexuscart/user-service/internal/server"
)

type serverRunner interface {
	Run(...string) error
}

var newServer = func() serverRunner {
	return server.New()
}

var execute = run

func main() {
	if err := execute(); err != nil {
		log.Fatalf("user-service stopped: %v", err)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	return newServer().Run(":" + port)
}