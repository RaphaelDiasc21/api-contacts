package main

import (
	"api-contacts/routes"
	"net/http"
	"github.com/gorilla/mux"
)



func main() {
	
	r := mux.NewRouter()

	routes.InitializeContactsRoutes(r)

	http.ListenAndServe(":8080",r)

}
