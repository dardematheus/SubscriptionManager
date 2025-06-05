package main

import (
	"log"
	"subscriptionmanager/internal/handlers"
	"subscriptionmanager/internal/routes"
	services "subscriptionmanager/internal/services"
)

func main() {
	db, err := services.StablishConnection()
	if err != nil {
		log.Fatalf("Failed to connect to DB")
	}
	defer db.Close()

	env := &handlers.Env{
		DB: db,
	}

	router := routes.InitRouter(env)

	log.Printf("Listening on port 1308")
	if err := router.Run(":1308"); err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
