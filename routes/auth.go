package routes

import (
	"golang-jwt/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST").Name("register")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
