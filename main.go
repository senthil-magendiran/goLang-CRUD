package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter()
	s.HandleFunc("/create", createUser).Methods("POST")
	s.HandleFunc("/getUsers", getUsers).Methods("GET")
	s.HandleFunc("/searchUser", searchUser).Methods("GET")
	s.HandleFunc("/updateUser", updateUser).Methods("PUT")
	s.HandleFunc("/deleteUser/{name}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s))
}
