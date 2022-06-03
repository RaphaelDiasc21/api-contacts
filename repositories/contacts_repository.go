package repositories

import (
	"api-contacts/database"
	"api-contacts/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ContactsRepository struct {
	databaseConnection *gorm.DB
}

func NewContactRepository(client string) *ContactsRepository {
	databaseConnection, err := database.GetDatabaseConnection(client)

	if err != nil {
		fmt.Println(err)
	}

	return &ContactsRepository{
		databaseConnection: databaseConnection,
	}
}

func (cr *ContactsRepository) InsertContact(contact *models.Contact) *gorm.DB {
	return cr.databaseConnection.Create(&contact)
}

func (cr *ContactsRepository) InsertContacts(contacts *[]models.Contact) *gorm.DB {
	return cr.databaseConnection.Omit("UpdatedAt","CreatedAt","DeletedAt").Create(contacts)
}

func (cr *ContactsRepository) FindAll() []models.Contact {
	contacts := []models.Contact{}
	cr.databaseConnection.Omit("UpdatedAt","CreatedAt","DeletedAt").Find(&contacts)
	return contacts
}

func (cr *ContactsRepository) FindById(contactId uint) (contact models.Contact, err error) {
	contact = models.Contact{}
	result := cr.databaseConnection.First(&contact,contactId)

	if result.RowsAffected == 0 {
		err = errors.New("Contact not found")
	}
	
	return
}

func (cr *ContactsRepository) Update(contactId uint,contact *models.Contact) (err error)  {
	contactFromDB := models.Contact{}
	result := cr.databaseConnection.First(&contactFromDB,contactId)

	if result.RowsAffected != 0 {
		contactFromDB.Nome = contact.Nome
		contactFromDB.Celular = contact.Celular
	
		cr.databaseConnection.Save(&contactFromDB)
		err = nil
	} else {
		err = errors.New("Contact Id not exists")
	}

	return
	
} 	


func (cr *ContactsRepository) Delete(contactId uint) (err error)  {
	contactFromDB := models.Contact{}
	result := cr.databaseConnection.First(&contactFromDB,contactId)

	if result.RowsAffected != 0 {
		cr.databaseConnection.Delete(&models.Contact{},contactId)
		err = nil
	} else {
		err = errors.New("Contact Id not exists")
	}

	return
} 	