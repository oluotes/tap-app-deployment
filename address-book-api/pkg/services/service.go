package service

import (
	"github.com/gin-gonic/gin"
	"vmware.com/migrator-poc/address-book-api/pkg/db"
	"vmware.com/migrator-poc/address-book-api/pkg/db/repository"
	"vmware.com/migrator-poc/address-book-api/pkg/models"
)

func GetContacts(c *gin.Context) []models.Contact {
	return repository.GetContacts(c)
}

func CreateContact(contact models.Contact, dbClient db.DBClient) uint {
	return repository.CreateContact(contact, dbClient)
}

func FindContact(id string, dbClient db.DBClient) (models.Contact, error) {
	contact, err := repository.FindContact(id, dbClient)
	if err != nil {
		return contact, err
	}
	return contact, nil
}

func DeleteContact(id string, dbClient db.DBClient) error {
	if err := repository.DeleteContact(id, dbClient); err != nil {
		return err
	}
	return nil

}
