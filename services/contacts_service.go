package services

import (
	"api-contacts/models"
	"api-contacts/repositories"
	"api-contacts/resources"
	"errors"
)

type ContactsService struct {
	contactRepository *repositories.ContactsRepository
	client string
}

func NewContactService(client string) *ContactsService {

	return &ContactsService{
		contactRepository: repositories.NewContactRepository(client),
		client: client,
	}

}

func (cs *ContactsService) InitializeContactModels(contactsResources *resources.ContactResources) (contactModels []models.Contact, contactModelsWithInvalidFormat []models.Contact) {
	for _,contactResource := range contactsResources.Contacts {
		contact, err := contactResource.ToModel(cs.client)

		if err != nil {
			contactModelsWithInvalidFormat = append(contactModelsWithInvalidFormat,*contact)
		} else {
			contactModels = append(contactModels,*contact)
		}
	}
	
	return
}

func (cs *ContactsService) ExtractCellphonesWithInvalidFormats(contactModelsWithInvalidFormat []models.Contact ) error {
	if len(contactModelsWithInvalidFormat) != 0 {
		constactStringToError := ""

		for _,contact := range contactModelsWithInvalidFormat {
			constactStringToError = constactStringToError + "Name: " + contact.Nome + " CellPhone: " + contact.Celular + ","
		}
		return errors.New("The followed contacts could not be saved on database because it is not on right standard:" + constactStringToError)
	} 

	return nil
}

func (cs *ContactsService) InsertContacts(contactResources *resources.ContactResources) error {
	contacts, contactsInvalid := cs.InitializeContactModels(contactResources)

	err := cs.ExtractCellphonesWithInvalidFormats(contactsInvalid)

	if err != nil {
		return err
	}

	cs.contactRepository.InsertContacts(&contacts)
	return nil
}

func (cs *ContactsService) FindAll() []resources.ContactResource {
	contacts := cs.contactRepository.FindAll()
	var contactResources []resources.ContactResource

	for _,resource := range contacts {
		contactResource := resources.ContactResource{
			ID: resource.ID,
			Name: resource.Nome,
			Cellphone: resource.Celular,
		}

		contactResources = append(contactResources,contactResource)
	}


	return contactResources
}

func (cs *ContactsService) FindById(contactId uint) (contactResource resources.ContactResource,err error) {
	contact, err := cs.contactRepository.FindById(contactId)

	if err == nil {
		contactResource = resources.ContactResource{
			ID: contact.ID,
			Name: contact.Nome,
			Cellphone: contact.Celular,
		}
	} 

	return
}

func (cs *ContactsService) Update(contactId uint, client string, contactResource resources.ContactResource) (err error) {
	contact, modelError := contactResource.ToModel(client)

	if modelError == nil {
		updateError := cs.contactRepository.Update(contactId,contact)

		if updateError != nil {
			err = updateError
		} else {
			err = nil
		}

		
	} else {
		err = modelError
	}
	return
}


func (cs *ContactsService) Delete(contactId uint, client string) (err error) {
	return cs.contactRepository.Delete(contactId)
}