package main

import (
	"log"

	"github.com/mx-ulises/go-learning/distributed-services-with-go/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
