package main

import (
	"log"
	"net/http"

	"testcodex/clients"
	"testcodex/controllers"
	"testcodex/routes"
	"testcodex/services"
)

func main() {
	client := clients.NewUserClient()
	service := services.NewUserService(client)
	controller := controllers.NewUserController(service)

	router := routes.SetupRoutes(controller)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
