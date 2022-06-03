package main

import (
	"api-contacts/routes"
	"net/http"
	"github.com/gorilla/mux"
)

// @title API Contacts
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name Raphael Dias
// @contact.email raphaeldiasc16@gmail.com
// @license.name Apache 2.0
// @host localhost:8080
// @BasePath /

func main() {
	
	r := mux.NewRouter()

	routes.InitializeContactsRoutes(r)

	http.ListenAndServe(":8080",r)

}
