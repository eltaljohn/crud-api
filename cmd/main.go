package main

import (
	"github.com/eltaljohn/crudapi/authorization"
	"github.com/eltaljohn/crudapi/handler"
	"github.com/eltaljohn/crudapi/storage"
	"log"
	"net/http"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Error loading certificates: %v", err)
	}
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Server running in port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("error to run serve %v", err)
	}
}
