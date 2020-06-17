package observation

import "github.com/vkevv/back-rectifier/pkg/models"

// Create creates a new observation
func (o *Observation) Create(documentID int, X, Y int, text, author string) (models.Observation, error) {
	observation := models.Observation{
		DocumentID: documentID,
		X:          X,
		Y:          Y,
		Text:       text,
		Author:     author,
	}
	err := o.DBActions.InsertObservation(&observation)
	return observation, err
}

// List get a list of observations of a document
func (o *Observation) List(documentID int) ([]models.Observation, error) {
	return o.DBActions.GetDocumentObs(documentID)
}

// Delete deletes a document
func (o *Observation) Delete(observationID int) error {
	return o.DBActions.DeleteObservation(observationID)
}
