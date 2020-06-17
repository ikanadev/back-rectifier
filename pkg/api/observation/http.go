package observation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/models"
	"github.com/vkevv/back-rectifier/pkg/server"
)

// HTTP is the main struct to handle observations request
type HTTP struct {
	service Service
}

// ServeHTTP takes a gin engine and attach handlers
func ServeHTTP(service Service, ginServer *gin.RouterGroup) {
	http := HTTP{service}
	ginServer.POST("/observation", http.PostObservation)
	ginServer.DELETE("/observation/:id", http.DeleteObservation)
	ginServer.GET("/document/:id/observations", http.ListObservations)
}

// PostObservation handler to register an obervation
func (h HTTP) PostObservation(c *gin.Context) {
	type reqVars struct {
		X          int    `json:"x"`
		Y          int    `json:"y"`
		Text       string `json:"text"`
		Author     string `json:"author"`
		DocumentID int    `json:"documentID"`
	}
	creds := reqVars{}
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	observation, err := h.service.Create(creds.DocumentID, creds.X, creds.Y, creds.Text, creds.Author)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, observation)
}

// DeleteObservation handler to delete an observation
func (h HTTP) DeleteObservation(c *gin.Context) {
	observationID, err := server.GetURIID(c)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	err = h.service.Delete(observationID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// ListObservations get a list of observations
func (h HTTP) ListObservations(c *gin.Context) {
	type response struct {
		Observations []models.Observation `json:"observations"`
	}
	documentID, err := server.GetURIID(c)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	observations, err := h.service.List(documentID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, response{observations})
}
