package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/models"
	"github.com/vkevv/back-rectifier/pkg/server"
)

// HTTP is the main struct to handle project request
type HTTP struct {
	service Service
}

// ServeHTTP takes a gin engine and sets the project related handlers
func ServeHTTP(service Service, ginServer *gin.RouterGroup) {
	http := HTTP{service}
	ginServer.POST("/project", http.PostProject)
	ginServer.GET("/project", http.ListProjects)
	ginServer.DELETE("/project/:id", http.DeleteProject)
}

// PostProject handler to register a project
func (h HTTP) PostProject(c *gin.Context) {
	type reqVars struct {
		Name      string `json:"name"`
		Institute string `json:"institute"`
	}
	creds := reqVars{}
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	project, err := h.service.Create(c, creds.Name, creds.Institute)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, project)
}

// ListProjects handler to register a project
func (h HTTP) ListProjects(c *gin.Context) {
	type response struct {
		Projects []models.Project `json:"projects"`
	}
	projects, err := h.service.List(c)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, response{projects})
}

// DeleteProject handler to register a project
func (h HTTP) DeleteProject(c *gin.Context) {
	type uriVars struct {
		ID string `uri:"id" binding:"required"`
	}
	uriParams := uriVars{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		server.ErrorResp(c, err)
		return
	}
	projectID, err := strconv.Atoi(uriParams.ID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	if err := h.service.Delete(projectID); err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
