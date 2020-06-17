package observation

import "github.com/vkevv/back-rectifier/pkg/models"

// Service all actions for observations
type Service interface {
	Create(documentID int, X, Y int, text, author string) (models.Observation, error)
	List(documentID int) ([]models.Observation, error)
	Delete(observationID int) error
}

// DBActions all db actions related to documents
type DBActions interface {
	InsertObservation(obs *models.Observation) error
	GetDocumentObs(documentID int) ([]models.Observation, error)
	DeleteObservation(obsID int) error
}

// Observation stuct which implements service
type Observation struct {
	DBActions
}

// LoadObservationService loads observation service
func LoadObservationService(dbActions DBActions) Observation {
	return Observation{dbActions}
}
