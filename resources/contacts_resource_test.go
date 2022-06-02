package resources_test

import (
	"api-contacts/resources"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestFormatCellPhoneByClientFailed(t *testing.T) {
	contactResource := resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "559875862152",
		}


	_, err := contactResource.FormatCellPhoneByClient("macapa",contactResource.Cellphone)

	assert.Equal(t,err.Error(),"cellphone is not in the right standard")
}


func TestMacapaFormatCellPhoneByClientSuccess(t *testing.T) {
	contactResource := resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "5598758621522",
		}


	cellPhone, err := contactResource.FormatCellPhoneByClient("macapa",contactResource.Cellphone)

	assert.Equal(t,err,nil)
	assert.Equal(t,cellPhone,"+55 (98) 75862-1522")
}

func TestVarejaoFormatCellPhoneByClientSuccess(t *testing.T) {
	contactResource := resources.ContactResource{
			ID: 1,
			Name: "Teste",
			Cellphone: "5598758621522",
		}


	cellPhone, err := contactResource.FormatCellPhoneByClient("varejao",contactResource.Cellphone)

	assert.Equal(t,err,nil)
	assert.Equal(t,cellPhone,"5598758621522")
}

func TestMacapaToModelSuccess(t *testing.T) {
	contactResource := resources.ContactResource{
		ID: 1,
		Name: "Teste",
		Cellphone: "5598758621522",
	}

	contactResources := resources.ContactResources{
		Contacts: contactResource,
	}

	contactResource.ToModel("macapa")
}