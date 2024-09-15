package main

import (
	"go-jwt/pkg/config"
	"go-jwt/pkg/controllers"
	"go-jwt/pkg/routes"
	"log"
	"net/http"
)

func main() {
	clientDb := config.ConnectDB("mongodb://localhost:27017")
	controllers.Collection = config.GetCollection(clientDb, "UserDb", "users")
	r := routes.NewAuthRoute()
	log.Println("Server running on port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
