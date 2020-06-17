package observation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/models"
	"github.com/vkevv/back-rectifier/pkg/server"
)

// HTTP is the main struct to handle observations request
type HTTP struct {
	service Service
}

// ServeHTTP takes a gin engine and attach handlers
func ServeHTTP(service Service, ginServer *gin.Engine) {
	http := HTTP{service}
	ginServer.GET("/v1/documentobs/:code", http.GetDocByCode)
	ginServer.POST("/v1/observation/:code", http.PostObservation)
	ginServer.DELETE("/v1/observation/:id/:code", http.DeleteObservation)
	ginServer.GET("/v1/document/:id/observations", http.ListObservations)
}

// GetDocByCode returns a document with all observations
func (h HTTP) GetDocByCode(c *gin.Context) {
	type uriVars struct {
		Code string `uri:"code" binding:"required"`
	}
	uriParams := uriVars{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		server.ErrorResp(c, err)
		return
	}
	document, err := h.service.GetDocByCode(uriParams.Code)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, document)
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
	type uriVars struct {
		Code string `uri:"code" binding:"required"`
	}
	uriParams := uriVars{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		server.ErrorResp(c, err)
		return
	}
	creds := reqVars{}
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	observation, err := h.service.Create(creds.DocumentID, creds.X, creds.Y, creds.Text, creds.Author, uriParams.Code)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, observation)
}

// DeleteObservation handler to delete an observation
func (h HTTP) DeleteObservation(c *gin.Context) {
	type uriVars struct {
		Code string `uri:"code" binding:"required"`
		ID   string `uri:"id" binding:"required"`
	}
	uriParams := uriVars{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		server.ErrorResp(c, err)
		return
	}
	observationID, err := strconv.Atoi(uriParams.ID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	err = h.service.Delete(observationID, uriParams.Code)
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
