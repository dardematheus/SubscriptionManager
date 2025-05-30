package main

import (
	"log"
	"subscriptionmanager/internal/routes"
	session "subscriptionmanager/internal/services"
)

func main() {
	db, err := session.StablishConnection()
	if err != nil {
		log.Fatalf("Failed to connect to DB")
	}
	defer db.Close()

	router := routes.InitRouter(db)

	log.Printf("Listening on port 1308")
	if err := router.Run(":1308"); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
