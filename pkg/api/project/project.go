package project

import (
	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/api/common"
	"github.com/vkevv/back-rectifier/pkg/models"
)

// Create creates a project
func (a *Project) Create(c *gin.Context, name, institute string) (models.Project, error) {
	project := models.Project{
		Name:      name,
		Institute: institute,
	}
	userID, exists := c.Get("id")
	if !exists {
		return project, common.ErrNoContextID
	}
	userIDInt, ok := userID.(int)
	if !ok {
		return project, common.ErrContextIDInvalid
	}
	project.UserID = userIDInt
	err := a.DBActions.InsertProject(&project)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}

// Delete deletes a project
func (a *Project) Delete(projectID int) error {
	project, err := a.DBActions.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	return a.DBActions.DeleteProject(project.ID)
}

// List list user projects
func (a *Project) List(c *gin.Context) ([]models.Project, error) {
	userID, exists := c.Get("id")
	if !exists {
		return nil, common.ErrNoContextID
	}
	userIDInt, ok := userID.(int)
	if !ok {
		return nil, common.ErrContextIDInvalid
	}
	projects, err := a.DBActions.GetUserProjects(userIDInt)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
