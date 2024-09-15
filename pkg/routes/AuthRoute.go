package routes

import (
	"go-jwt/pkg/controllers"

	"github.com/gorilla/mux"
)

func NewAuthRoute() *mux.Router {
	// Create a new router
	router := mux.NewRouter()
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	return router
}
