package main

import (
	"SubscriptionManager/internal/routes"
	"log"
)

func main() {
	router := routes.InitRouter()

	log.Printf("Listening on port 1308")
	if err := router.Run(":1308"); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
