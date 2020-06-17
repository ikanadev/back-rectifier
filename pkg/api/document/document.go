package document

import (
	"io"

	"github.com/vkevv/back-rectifier/pkg/models"
)

// Create creates and stores a new document
func (d *Document) Create(projectID int, comment, fileName string, reader io.Reader) (models.Document, error) {
	document := models.Document{
		Comment:   comment,
		ProjectID: projectID,
		Rectified: false,
	}
	err := document.GenerateRandomCode()
	if err != nil {
		return document, err
	}
	exists, err := d.DBActions.ExistsDocCode(document.AccessCode)
	if err != nil {
		return document, err
	}
	for exists {
		err = document.GenerateRandomCode()
		if err != nil {
			return document, err
		}
		exists, err = d.DBActions.ExistsDocCode(document.AccessCode)
		if err != nil {
			return document, err
		}
	}
	url, err := d.FileActions.UploadFile(fileName, reader)
	if err != nil {
		return document, err
	}
	document.URL = url
	err = d.DBActions.InsertDocument(&document)
	if err != nil {
		return document, err
	}
	return document, nil
}

// GenerateCode Generates a new code for a document
func (d *Document) GenerateCode(documentID int) (models.Document, error) {
	document, err := d.DBActions.GetDocumentByID(documentID)
	if err != nil {
		return models.Document{}, err
	}
	err = document.GenerateRandomCode()
	if err != nil {
		return models.Document{}, err
	}
	err = d.DBActions.UpdateDocument(&document)
	if err != nil {
		return models.Document{}, err
	}
	return document, nil
}

// Delete deletes a document
func (d *Document) Delete(documentID int) error {
	return d.DBActions.DeleteDocument(documentID)
}

// List get all documents of a project
func (d *Document) List(projectID int) ([]models.Document, error) {
	return d.DBActions.GetProjectDocuments(projectID)
}
