package common

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/vkevv/back-rectifier/pkg/models"
)

// DB all database actions
type DB struct {
	db orm.DB
}

// NewDB creates a new instance
func NewDB(db orm.DB) DB {
	return DB{db}
}

// GetUserByID get by ID
func (d DB) GetUserByID(id int) (models.User, error) {
	user := models.User{}
	user.ID = id
	err := d.db.Select(&user)
	return user, err
}

// GetUserByEmail get user by email
func (d DB) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	user.Email = email
	err := d.db.Model(&user).Where("email = ?email").First()
	return user, err
}

// ExistsEmail check email existence
func (d DB) ExistsEmail(email string) (bool, error) {
	user := models.User{}
	user.Email = email
	return d.db.Model(&user).Where("email = ?email").Exists()
}

// InsertUser insert user in db
func (d DB) InsertUser(user *models.User) error {
	return d.db.Insert(user)
}

// InsertProject stores a project
func (d DB) InsertProject(project *models.Project) error {
	return d.db.Insert(project)
}

// GetUserProjects get projects of an user
func (d DB) GetUserProjects(userID int) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	err := d.db.Model(&projects).Where("user_id = ?", userID).Order("id ASC").Select()
	return projects, err
}

// DeleteProject soft del a project
func (d DB) DeleteProject(projectID int) error {
	project := models.Project{}
	project.ID = projectID
	err := d.db.Delete(&project)
	return err
}

// GetProjectByID get a project by id
func (d DB) GetProjectByID(projectID int) (models.Project, error) {
	project := models.Project{}
	project.ID = projectID
	err := d.db.Select(&project)
	return project, err
}

// InsertDocument insert a document in database
func (d DB) InsertDocument(document *models.Document) error {
	return d.db.Insert(document)
}

// GetProjectDocuments get all documents of a project
func (d DB) GetProjectDocuments(projectID int) ([]models.Document, error) {
	documents := make([]models.Document, 0)
	err := d.db.Model(&documents).Where("project_id = ?", projectID).Order("id ASC").Select()
	return documents, err
}

// DeleteDocument deletes a document given an id
func (d DB) DeleteDocument(documentID int) error {
	document := models.Document{}
	document.ID = documentID
	return d.db.Delete(&document)
}

// UpdateDocument updates data of a document (id field must be valid)
func (d DB) UpdateDocument(document *models.Document) error {
	return d.db.Update(document)
}

// GetDocumentByID get a document by ID
func (d DB) GetDocumentByID(documentID int) (models.Document, error) {
	document := models.Document{}
	document.ID = documentID
	err := d.db.Select(&document)
	return document, err
}

// InsertObservation insert an observation
func (d DB) InsertObservation(obs *models.Observation) error {
	return d.db.Insert(obs)
}

// GetDocumentObs get all observations of a document
func (d DB) GetDocumentObs(documentID int) ([]models.Observation, error) {
	observations := make([]models.Observation, 0)
	err := d.db.Model(&observations).Where("document_id = ?", documentID).Order("y ASC").Select()
	return observations, err
}

// DeleteObservation deletes an observation
func (d DB) DeleteObservation(obsID int) error {
	observation := models.Observation{}
	observation.ID = obsID
	return d.db.Delete(&observation)
}
