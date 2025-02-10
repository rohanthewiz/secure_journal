package main

import (
	"log"
	"secure_journal/web"
)

func main() {
	s := web.InitWeb()
	log.Println(s.Run())
}
