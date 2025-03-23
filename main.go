package main

import (
	"log"
	"secure_journal/login"
	"secure_journal/web"
)

func main() {
	err := login.InitDB("users.duck")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer login.CloseDB()

	s := web.InitWeb()
	log.Println(s.Run())
}
