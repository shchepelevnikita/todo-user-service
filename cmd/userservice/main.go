package main

import (
	"log"
	"net/http"

	"github.com/shchepelevnikita/todo-user-service/internal/userservice"
)

func main() {
	router := userservice.SetupRouter()
	log.Println("Starting User Service on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
