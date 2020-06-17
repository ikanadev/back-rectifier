package models

// Project a project contains 1 or many documents
type Project struct {
	Base
	Name      string      `json:"name"`
	Institute string      `json:"institute"`
	UserID    int         `json:"userID"`
	User      *User       `json:"user"`
	Documents []*Document `json:"documents"`
}
