package routes

import (
	a "CrudSystem/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter() // created a router

	// implementing different routes
	r.HandleFunc("/users", a.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", a.GetUser).Methods("GET")
	r.HandleFunc("/users", a.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", a.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", a.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r)) // (port, router)
}
