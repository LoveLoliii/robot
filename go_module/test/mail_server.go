package main

import (
	"log"

	"github.com/emersion/go-imap/backend/memory"
	"github.com/emersion/go-imap/server"
)

func main() {
	// Create a memory backend
	be := memory.New()

	//Create a new server
	s := server.New(be)
	s.Addr = ":1143"

	s.AllowInsecureAuth = true

	log.Println("Starting IMAP server at localhost:1143")
	if err := s.ListenAndServer(); err != nil {
		log.Fatal(err)
	}
}
