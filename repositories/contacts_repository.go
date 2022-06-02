package repositories

import (
	"api-contacts/database"
	"api-contacts/models"
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