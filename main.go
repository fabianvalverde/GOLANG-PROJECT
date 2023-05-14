package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	User_Id          int    `json:"user_id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Shipping_address string `json:"shipping_address"`
}

type Users []User

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{User_Id: 123, Name: "Fabian", Email: "fabian3830@gmail.com", Password: "1245", Shipping_address: "fdgdfgd"},
	}

	fmt.Println("Endpoint Hit: All users Endpoint")
	json.NewEncoder(w).Encode(users)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", allUsers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
