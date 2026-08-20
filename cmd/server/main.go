package main

import (
	"log"
	"os"

	"nexuscart/user-service/internal/server"
)

type serverRunner interface {
	Run(...string) error
}

func defaultNewServer() serverRunner {
	return server.New()
}

var newServer = defaultNewServer
var fatalf = log.Fatalf
var execute = run

func main() {
	if err := execute(); err != nil {
		fatalf("user-service stopped: %v", err)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	return newServer().Run(":" + port)
}
