package models

// Observation represents an document observation
type Observation struct {
	Base
	X          int       `json:"x"`
	Y          int       `json:"y"`
	Text       string    `json:"text"`
	Author     string    `json:"author"`
	Document   *Document `json:"document"`
	DocumentID int       `json:"documentID"`
}
