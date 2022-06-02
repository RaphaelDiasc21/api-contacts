package services_test

import (
	"api-contacts/resources"
	"api-contacts/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMacapaContactResourceToModelSucess(t *testing.T) {
	contactResource := []resources.ContactResource{
		resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "5598758621521",
		},
	}


	contactResources := resources.ContactResources{
		Contacts: contactResource,
	}

	cs := services.NewContactService("macapa")
	contactsModel,contactWithInvalidFormat := cs.InitializeContactModels(&contactResources)

	assert.Equal(t,contactsModel[0].Celular,"+55 (98) 75862-1521")
	assert.Equal(t,len(contactWithInvalidFormat),0)
}

func TestVarejaoContactResourceToModelSucess(t *testing.T) {
	contactResource := []resources.ContactResource{
		resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "5598758621521",
		},
	}


	contactResources := resources.ContactResources{
		Contacts: contactResource,
	}

	cs := services.NewContactService("varejao")
	contactsModel,contactWithInvalidFormat := cs.InitializeContactModels(&contactResources)

	assert.Equal(t,contactsModel[0].Celular,"5598758621521")
	assert.Equal(t,len(contactWithInvalidFormat),0)
}


func TestVarejaoContactWithInvalidFormat(t *testing.T) {
	contactResource := []resources.ContactResource{
		resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "55987586215210",
		},
	}


	contactResources := resources.ContactResources{
		Contacts: contactResource,
	}

	cs := services.NewContactService("varejao")
	contactsModel,contactWithInvalidFormat := cs.InitializeContactModels(&contactResources)

	assert.Equal(t,len(contactsModel),0)
	assert.Equal(t,len(contactWithInvalidFormat),1)
}

func TestExtractCellPhonesWithInvalidFormat(t *testing.T) {
	contactResource := []resources.ContactResource{
		resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "55987586215210",
		},
	}


	contactResources := resources.ContactResources{
		Contacts: contactResource,
	}

	cs := services.NewContactService("varejao")
	_,contactWithInvalidFormat := cs.InitializeContactModels(&contactResources)

	err := cs.ExtractCellphonesWithInvalidFormats(contactWithInvalidFormat)

	assert.NotEqual(t,err,nil)

}