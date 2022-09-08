package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vmware.com/migrator-poc/address-book-api/pkg/models"
)

func InitDB() (*gorm.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	hostPort := os.Getenv("DB_HOST_PORT")
	connectionStr := username + ":" + password + "@(" + hostname + ":" + hostPort + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	log.Println(connectionStr)
	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		//log.Panic(err)
		log.Fatalln(err)
		// return nil, err
	}
	db.AutoMigrate(&models.Contact{})

	return db, nil
}
