package routes

import (
	"golang-jwt/controllers"
	"golang-jwt/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
