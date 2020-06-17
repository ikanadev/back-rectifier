package document

import (
	"io"

	"github.com/vkevv/back-rectifier/pkg/models"
)

// Service all handler for documents
type Service interface {
	Create(projectID int, comment, fileName string, reader io.Reader) (models.Document, error)
	GenerateCode(documentID int) (models.Document, error)
	Delete(documentID int) error
	List(projectID int) ([]models.Document, error)
}

// DBActions direct database actions needed
type DBActions interface {
	InsertDocument(document *models.Document) error
	GetProjectDocuments(projectID int) ([]models.Document, error)
	DeleteDocument(documentID int) error
	GetDocumentByID(documentID int) (models.Document, error)
	UpdateDocument(document *models.Document) error
}

// FileActions an interface to represent file actions
type FileActions interface {
	UploadFile(fileName string, file io.Reader) (string, error)
}

// Document struct which implements service
type Document struct {
	DBActions
	FileActions
}

// LoadDocumentService loads document service
func LoadDocumentService(dbActions DBActions, fileActions FileActions) Document {
	return Document{dbActions, fileActions}
}
