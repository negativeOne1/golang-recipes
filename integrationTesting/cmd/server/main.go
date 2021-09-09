package main

import (
	"httpServerBareMetal/internal/server"
	"log"
)

const (
	httpAddress = ":8080"
)

func main() {
	if err := r(); err != nil {
		log.Fatal(err)
	}
}

func r() error {
	s := server.New()

	return s.ListenAndServe(httpAddress)
}
