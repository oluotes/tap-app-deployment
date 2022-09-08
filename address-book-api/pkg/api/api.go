package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vmware.com/migrator-poc/address-book-api/pkg/db"
	"vmware.com/migrator-poc/address-book-api/pkg/models"
	service "vmware.com/migrator-poc/address-book-api/pkg/services"
)

var addressBook = []models.Contact{
	{
		FirstName:    "Olumide",
		LastName:     "Oteniya",
		Street:       "35 Hunterview Drive",
		City:         "Calgary",
		Country:      "Canada",
		PostalCode:   "T3P0 W7",
		Province:     "Alberta",
		Email:        "oolumide@vmware.com",
		OfficeNumber: "+1 315 333 3000",
		HomeNumber:   "+1 487 532 3501",
		MobileNumber: "+1 456 322 3456",
	},
}

// GetContacts godoc
// @Summary      Get all Contacts
// @Description  Get all Contacts
// @Tags         GetContacts
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string][]models.Contact
// @Router       /contacts [get]
func GetContacts(c *gin.Context) {
	addressBook := service.GetContacts(c)
	c.IndentedJSON(http.StatusOK, gin.H{"data": addressBook})
}

// FindContact godoc
// @Summary      Find a Contact
// @Description  get contact by ID
// @Tags         FindContact
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Contact ID"
// @Success      200  {object}  models.Contact
// @Failure      400  {object}  map[string]string
// @Router       /contacts/{id} [get]
func FindContact(c *gin.Context) {
	dbClient := c.MustGet("dbConn").(db.DBClient)
	id := c.Param("id")

	contact, err := service.FindContact(id, dbClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, contact)
}

// CreateContact godoc
// @Summary      Create a Contact
// @Description  get contact by ID
// @Tags         CreateContact
// @Accept       json
// @Produce      json
// @Param        message  body      models.Contact  true  "Contact"
// @Success      200  {object}  map[string]integer
// @Failure      400  {object}  map[string]string
// @Router       /contacts/create [post]
func CreateContact(c *gin.Context) {
	dbClient := c.MustGet("dbConn").(db.DBClient)
	var contact models.Contact
	err := c.ShouldBindJSON(&contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ID": service.CreateContact(contact, dbClient)})
}

// UpdateContact godoc
// @Summary      Update a Contact
// @Description  Update contact by ID
// @Tags         UpdateContact
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Contact ID"
// @Param        message  body      models.Contact  true  "Contact"
// @Success      200  {object}  models.Contact
// @Failure      400  {object}  map[string]string
// @Router       /contacts/{id} [patch]
func UpdateContact(c *gin.Context) {
	dbClient := c.MustGet("dbConn").(db.DBClient)
	id := c.Param("id")
	contact, err := service.FindContact(id, dbClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var inputContact models.Contact
	if err := c.ShouldBindJSON(&inputContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbClient.Db.Model(&contact).Updates(inputContact)

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// DeleteContact godoc
// @Summary      Delete a Contact
// @Description  delete contact by ID
// @Tags         DeleteContact
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Contact ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /contacts/{id} [delete]
func DeleteContact(c *gin.Context) {
	dbClient := c.MustGet("dbConn").(db.DBClient)
	id := c.Param("id")

	if err := service.DeleteContact(id, dbClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "Deleted"})
}
