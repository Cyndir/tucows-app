package main

import (
	"log"

	"github.com/Cyndir/tucows-app/internal/database"
	"github.com/Cyndir/tucows-app/internal/messagequeue"
	"github.com/Cyndir/tucows-app/internal/router"
)

func main() {
	router := router.New(database.New(), messagequeue.New())
	router.Setup()
	log.Println("Setup complete, listening on port 8080")
	router.Run()
	log.Println("Server closing on port 8080")
}
