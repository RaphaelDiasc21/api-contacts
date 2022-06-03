package routes

import (
	"api-contacts/handlers"
	"github.com/gorilla/mux"
)
var spec []byte

func InitializeContactsRoutes(router *mux.Router) {
	router.HandleFunc("/contacts/{client}",handlers.CreateContactHandler).Methods("POST")
	router.HandleFunc("/contacts/{client}",handlers.FindAllHander).Methods("GET")
	router.HandleFunc("/contacts/{client}/{id}",handlers.FindByIdHandler).Methods("GET")
	router.HandleFunc("/contacts/{client}/{id}",handlers.UpdateHandler).Methods("PUT")
	router.HandleFunc("/contacts/{client}/{id}",handlers.DeleteHandler).Methods("DELETE")
}