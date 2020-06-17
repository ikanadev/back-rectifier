package observation

import (
	"github.com/vkevv/back-rectifier/pkg"
	"github.com/vkevv/back-rectifier/pkg/models"
)

const (
	noValidCode = "NOT_VALID_CODE"
)

var (
	errNoValidCode = pkg.NewAPIErr("Provided code is invalid or no longer valid", noValidCode)
)

// Create creates a new observation
func (o *Observation) Create(documentID int, X, Y int, text, author, code string) (models.Observation, error) {
	document, err := o.DBActions.GetDocumentByID(documentID)
	if err != nil {
		return models.Observation{}, err
	}
	if document.AccessCode != code {
		return models.Observation{}, errNoValidCode
	}
	observation := models.Observation{
		DocumentID: documentID,
		X:          X,
		Y:          Y,
		Text:       text,
		Author:     author,
	}
	err = o.DBActions.InsertObservation(&observation)
	return observation, err
}

// List get a list of observations of a document
func (o *Observation) List(documentID int) ([]models.Observation, error) {
	return o.DBActions.GetDocumentObs(documentID)
}

// Delete deletes a document
func (o *Observation) Delete(observationID int, code string) error {
	document, err := o.DBActions.GetDocumentByCode(code)
	if err != nil {
		return err
	}
	observations, err := o.DBActions.GetDocumentObs(document.ID)
	if err != nil {
		return err
	}
	existObsID := false
	for _, obs := range observations {
		if obs.ID == observationID {
			existObsID = true
			break
		}
	}
	if !existObsID {
		return errNoValidCode
	}
	return o.DBActions.DeleteObservation(observationID)
}

// GetDocByCode returns a document with all observations
func (o *Observation) GetDocByCode(code string) (models.Document, error) {
	document, err := o.DBActions.GetDocumentByCode(code)
	if err != nil {
		return models.Document{}, err
	}
	observations, err := o.DBActions.GetDocumentObs(document.ID)
	if err != nil {
		return models.Document{}, err
	}
	document.Observations = make([]*models.Observation, 0)
	for _, observation := range observations {
		document.Observations = append(document.Observations, &observation)
	}
	return document, nil
}
