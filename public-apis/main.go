package main

import (
	"log"
	"net/http"
	"public-apis/clients"
	"public-apis/handlers"
	"public-apis/utils"
	"time"
)

func main() {
	newReq := utils.NewRequest(10 * time.Second)
	userService := clients.NewUserService("http://localhost:8080/users", newReq)
	listingService := clients.NewListingService("http://localhost:6000/listings", newReq)
	restHandler := handlers.NewRestHandler(userService, listingService)

	// Set up routes
	http.HandleFunc("/public-api/users", restHandler.CreateUser)
	http.HandleFunc("/public-api/listings", restHandler.CreateOrGetListings)

	log.Println("Server starting on :8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
