package models

// Project a project contains 1 or many documents
type Project struct {
	Base
	Name      string `json:"name"`
	Institute string `json:"institute"`
	User      *User  `json:"user"`
	UserID    int
}
