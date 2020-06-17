package project

import (
	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/models"
)

// Service all handlers for project
type Service interface {
	Create(c *gin.Context, name, institute string) (models.Project, error)
	Delete(projecID int) error
	List(c *gin.Context) ([]models.Project, error)
}

// LoadProjectService loads service
func LoadProjectService(dbActions DBActions) Project {
	return Project{
		DBActions: dbActions,
	}
}

// Project implements Service
type Project struct {
	DBActions DBActions
}

// DBActions direct database actions
type DBActions interface {
	InsertProject(project *models.Project) error
	GetUserProjects(userID int) ([]models.Project, error)
	DeleteProject(projectID int) error
	GetProjectByID(projectID int) (models.Project, error)
}
