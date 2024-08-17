package main

import (
    "encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type User struct {
	ID int
	NAME string
	EMAIL string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	users := []User{}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", GetUser).Methods("GET")

    fmt.Println("Servidor rodando na porta 8000")
    http.ListenAndServe(":8000", router)
}