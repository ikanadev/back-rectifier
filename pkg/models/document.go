package models

import (
	"crypto/rand"
	"fmt"
)

// Document repretents a document to be revised
type Document struct {
	Base
	URL          string         `json:"url"`
	Comment      string         `json:"comment"`
	Rectified    bool           `json:"rectified"`
	AccessCode   string         `json:"accessCode"`
	Project      *Project       `json:"project"`
	ProjectID    int            `json:"projectID"`
	Observations []*Observation `json:"observations"`
}

// GenerateRandomCode generates a random code to access to the document
func (d *Document) GenerateRandomCode() error {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	s := fmt.Sprintf("%X-%X-%X", b[:2], b[2:4], b[4:])
	d.AccessCode = s
	return nil
}
