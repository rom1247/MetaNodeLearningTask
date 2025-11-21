package main

import (
	"backend/backend/phase_two/task_four/internal/bootstrap"
	"log"
)

// @
func main() {
	app, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
