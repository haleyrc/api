package main

import (
	"log"
	"net/http"

	"github.com/haleyrc/api/app"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var server app.Server
	log.Println("listening on :8080...")
	return http.ListenAndServe(":8080", &server)
}
