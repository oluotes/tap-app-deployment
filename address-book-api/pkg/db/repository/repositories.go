package repository

import (
	"github.com/gin-gonic/gin"
	"vmware.com/migrator-poc/address-book-api/pkg/db"
	"vmware.com/migrator-poc/address-book-api/pkg/models"
)

func GetContacts(c *gin.Context) []models.Contact {
	// db := c.MustGet("dbConn").(*gorm.DB)
	dbClient := c.MustGet("dbConn").(db.DBClient)
	db := dbClient.Db
	var addressBook []models.Contact
	db.Find(&addressBook)
	return addressBook
}

func FindContact(id string, dbClient db.DBClient) (models.Contact, error) {
	var contact models.Contact
	err := dbClient.Db.Where("id = ?", id).First(&contact).Error

	if err != nil {
		return contact, err
	}
	return contact, nil
}

func CreateContact(contact models.Contact, dbClient db.DBClient) uint {
	dbClient.Db.Save(&contact)
	return contact.ID
}

func DeleteContact(id string, dbClient db.DBClient) error {
	var contact models.Contact
	err := dbClient.Db.Where("id = ?", id).First(&contact).Error

	if err != nil {
		return err
	}
	dbClient.Db.Delete(&contact)
	return nil
}
