package models

// type ContactDetail struct {
// 	gorm.Model
// 	FirstName string
// 	LastName  string
// 	AddressID int
// 	Address   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	ContactID int
// 	Contact   Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// type Contact struct {
// 	ID            int
// 	Email         string
// 	PhoneNumberID int
// 	PhoneNumber   PhoneNumber `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// type PhoneNumber struct {
// 	ID     int
// 	Office string
// 	Home   string
// 	Mobile string
// }

// type Address struct {
// 	ID         int
// 	Street     string
// 	City       string
// 	Province   string
// 	Country    string
// 	PostalCode string
// }
type Contact struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Street       string `json:"street"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	PostalCode   string `json:"postalCode"`
	Email        string `json:"email"`
	OfficeNumber string `json:"work"`
	HomeNumber   string `json:"home"`
	MobileNumber string `json:"mobile"`
}
