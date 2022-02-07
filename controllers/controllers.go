package controllers

import (
	a "CrudSystem/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// handler methods
// http.ResponseWriter -> sends data to the HTTP client
// http.Request -> data structure that represent a client HTTP request
// *http.Request -> pointer of the HTTP request
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users [](a.User)
	a.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user [](a.User)
	a.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user [](a.User)
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("h1")
	a.DB.Create(&user)
	fmt.Println("h2")
	json.NewEncoder(w).Encode(user)
	fmt.Println("h3")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user [](a.User)
	a.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	a.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user [](a.User)
	a.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
