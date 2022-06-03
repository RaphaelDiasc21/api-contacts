package resources

import (
	"api-contacts/models"
	"errors"
	"regexp"
	"strings"
)

type ContactResource struct {
	ID uint `json:"_id"`
	Name string `json:name`
	Cellphone string `json:cellphone`
}

type ContactResources struct {
	Contacts []ContactResource `json:contacts`
}

func (cr *ContactResource) FormatCellPhoneByClient(client string, cellPhone string) (cellPhoneFormated string , err error) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	cellPhoneFormated = strings.Trim(cellPhone," ")
	cellPhoneFormated = reg.ReplaceAllString(cellPhoneFormated,"")

	if len(cellPhoneFormated) != 13 {
		err = errors.New("cellphone is not in the right standard")
		return
	}

	if client == "macapa" {
		cellphoneSplited := strings.Split(cellPhoneFormated,"")
		cellPhoneFormated = "+" + cellphoneSplited[0] + cellphoneSplited[1] + " (" + cellphoneSplited[2] + cellphoneSplited[3] + ") " + strings.Join(cellphoneSplited[4:9],"") + "-" + strings.Join(cellphoneSplited[9:],"")
		return
	}

	return
}

func (cr *ContactResource) ToModel(client string) (*models.Contact, error) {

	cellPhone,err := cr.FormatCellPhoneByClient(client,cr.Cellphone)

	
	return &models.Contact{
		ID: cr.ID,
		Nome: cr.Name,
		Celular: cellPhone,
	}, err
}

func (cr *ContactResource) ToResource(contact models.Contact) *ContactResource {
	return &ContactResource{
		ID: contact.ID,
		Name: contact.Nome,
		Cellphone: contact.Celular,
	}
}