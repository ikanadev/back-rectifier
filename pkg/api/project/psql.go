package project

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/vkevv/back-rectifier/pkg/models"
)

// DB database actions
type DB struct{}

// InsertProject stores a project
func (a DB) InsertProject(db orm.DB, project *models.Project) error {
	return db.Insert(project)
}

// GetUserProjects get projects of an user
func (a DB) GetUserProjects(db orm.DB, userID int) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	err := db.Model(&projects).Where("user_id = ?", userID).Order("id ASC").Select()
	return projects, err
}

// DeleteProject soft del a project
func (a DB) DeleteProject(db orm.DB, projectID int) error {
	project := models.Project{}
	project.ID = projectID
	return db.Delete(project)
}

// GetProjectByID get a project by id
func (a DB) GetProjectByID(db orm.DB, projectID int) (models.Project, error) {
	project := models.Project{}
	project.ID = projectID
	err := db.Select(&project)
	return project, err
}
