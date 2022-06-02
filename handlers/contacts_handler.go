package handlers

import (
	"api-contacts/resources"
	"api-contacts/services"
	//"api-contacts/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	
	var contactResources resources.ContactResources
	
	err := json.NewDecoder(r.Body).Decode(&contactResources)

	if err != nil {
		fmt.Println(err)
	}

	vars := mux.Vars(r)
	client,ok := vars["client"]

	if !ok {
		fmt.Println("Client is missing")
	}

	contactsService := services.NewContactService(client)
	err = contactsService.InsertContacts(&contactResources)

	w.Header().Set("Content-Type","application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse,e := json.Marshal(map[string]string{"message":err.Error()})
		fmt.Println(e)

		w.Write(jsonResponse)
	} else {
		jsonResponse,e := json.Marshal(map[string]string{"message":"Contacts saved !"})
		w.WriteHeader(http.StatusOK)
		fmt.Println(e)
		w.Write(jsonResponse)
	}

}