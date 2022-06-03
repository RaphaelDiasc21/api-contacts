package handlers

import (
	"api-contacts/resources"
	"api-contacts/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ExtractClientFromPath(r *http.Request,path string) (string, bool){
	vars := mux.Vars(r)
	client,ok := vars[path]

	return client,ok
}


func WriteJsonResponse(w http.ResponseWriter,httpStatus int, json []byte) {
	w.WriteHeader(httpStatus)
	w.Write(json)
}

// @Summary Save contacts
// @Description Save contacts on database
// @Accept  json
// @Produce  json
// @Success 201
// @Router /contacts/{client} [post]
func CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var contactResources resources.ContactResources
	
	err := json.NewDecoder(r.Body).Decode(&contactResources)

	if err == nil {
		client,_ := ExtractClientFromPath(r,"client")
		contactsService := services.NewContactService(client)
		err = contactsService.InsertContacts(&contactResources)
	
		if err != nil {
			json,_ := json.Marshal(map[string]string{"message":err.Error()})
			WriteJsonResponse(w,http.StatusBadRequest,json)
		} else {
			json,_ := json.Marshal(map[string]string{"message":"Contacts saved !"})
			WriteJsonResponse(w,http.StatusCreated,json)
		}
	} else {
		json,_ := json.Marshal(map[string]string{"message":"Error parsing body"})
		WriteJsonResponse(w,http.StatusBadRequest,json)
	}
}

// @Summary Find contacts
// @Description Find contacts on database
// @Accept  json
// @Produce  json
// @Success 200
// @Router /contacts/{client} [get]
func FindAllHander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	client,_ := ExtractClientFromPath(r,"client")

	contactsService := services.NewContactService(client)
	contactResources := contactsService.FindAll()

	json,_ := json.Marshal(contactResources)
	WriteJsonResponse(w,http.StatusOK,json)
}

// @Summary Find contact by id
// @Description Find contact by id on database
// @Accept  json
// @Produce  json
// @Success 200
// @Router /contacts/{client}/{id} [get]
func FindByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	client,_ := ExtractClientFromPath(r,"client")
	id,_ := ExtractClientFromPath(r,"id")

	contactsService := services.NewContactService(client)

	idInteger, _ := strconv.Atoi(id)
	contactResource,err := contactsService.FindById(uint(idInteger))

	if err == nil {
		json,_ := json.Marshal(contactResource)
		WriteJsonResponse(w,http.StatusOK,json)
	} else {
		json,_ := json.Marshal(err.Error())
		WriteJsonResponse(w,http.StatusNotFound,json)
	}

}

// @Summary Update contact by id
// @Description Update contact by id on database
// @Accept  json
// @Produce  json
// @Success 204
// @Error 404
// @Router /contacts/{client}/{id} [put]
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	client,_ := ExtractClientFromPath(r,"client")
	id,_ := ExtractClientFromPath(r,"id")

	var contactResource resources.ContactResource
	
	err := json.NewDecoder(r.Body).Decode(&contactResource)

	if err == nil {
		contactsService := services.NewContactService(client)

		idInteger, _ := strconv.Atoi(id)
		err = contactsService.Update(uint(idInteger),client,contactResource)

		if err == nil {
			json,_ := json.Marshal(map[string]string{"message":"Contact upated"})
			WriteJsonResponse(w,204,json)
		} else {
			json,_ := json.Marshal(err.Error())
			WriteJsonResponse(w,http.StatusNotFound,json)
		}

	} else {
		json,_ := json.Marshal(map[string]string{"message":"Error parsing body"})
		WriteJsonResponse(w,http.StatusBadRequest,json)
	}

}


// @Summary Delete contact by id
// @Description Delete contact by id on database
// @Accept  json
// @Produce  json
// @Success 200
// @Router /contacts/{client}/{id} [delete]
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	client,_ := ExtractClientFromPath(r,"client")
	id,_ := ExtractClientFromPath(r,"id")

	idInteger, _ := strconv.Atoi(id)
	contactsService := services.NewContactService(client)
	err := contactsService.Delete(uint(idInteger),client)

	if err == nil {
		json,_ := json.Marshal(map[string]string{"message":"Contact deleted"})
		WriteJsonResponse(w,http.StatusOK,json)

	} else {
		json,_ := json.Marshal(err.Error())
		WriteJsonResponse(w,http.StatusBadRequest,json)
	}

}