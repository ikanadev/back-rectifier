package models

// User who stores projects and documents
type User struct {
	Base
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
