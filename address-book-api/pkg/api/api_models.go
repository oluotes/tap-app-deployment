package api

// type ContactDetail struct {
// 	FirstName string  `json:"firstname"`
// 	LastName  string  `json:"lastname"`
// 	Address   Address `json:"address"`
// 	Contact   Contact `jsopn:"contact"`
// }

// type Contact struct {
// 	Email       string      `json:"email"`
// 	PhoneNumber PhoneNumber `json:"phone"`
// }

// type PhoneNumber struct {
// 	Office string `json:"work"`
// 	Home   string `json:"home"`
// 	Mobile string `json:"mobile"`
// }

// type Address struct {
// 	Street     string `json:"street"`
// 	City       string `json:"city"`
// 	Province   string `json:"province"`
// 	Country    string `json:"country"`
// 	PostalCode string `json:"postalCode"`
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
