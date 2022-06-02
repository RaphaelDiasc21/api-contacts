package routes

import (
	"api-contacts/handlers"
	"github.com/gorilla/mux"
)

func InitializeContactsRoutes(router *mux.Router) {
	router.HandleFunc("/contacts/{client}",handlers.CreateContactHandler).Methods("POST")
}