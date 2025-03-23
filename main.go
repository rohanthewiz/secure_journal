package main

import (
	"log"
	"secure_journal/login"
	"secure_journal/web"
)

func main() {
	// Initialize the database and get the connection
	db, err := login.InitDB("users.duck")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close() // Use db.Close() instead of login.CloseDB()

	// Pass the database to InitWeb
	s := web.InitWeb(db)

	// Start the server
	log.Println(s.Run())
}
