package document

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/models"
	"github.com/vkevv/back-rectifier/pkg/server"
)

// HTTP is the main struct to handle document request
type HTTP struct {
	service Service
}

// PostDocument handler to register a document, this only accepts multipart/form-data
func (h HTTP) PostDocument(c *gin.Context) {
	projectIDStr := c.PostForm("projectId")
	comment := c.PostForm("comment")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	fileHeader, err := c.FormFile("document")
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	newFileName := strconv.Itoa(int(fileHeader.Size)) + strings.ReplaceAll(fileHeader.Filename, " ", "_")
	document, err := h.service.Create(projectID, comment, newFileName, file)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, document)
}

// ListDocuments handler to list all documents of a project
func (h HTTP) ListDocuments(c *gin.Context) {
	type uriVars struct {
		ID string `uri:"id" binding:"required"`
	}
	type response struct {
		Documents []models.Document `json:"documents"`
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
	documents, err := h.service.List(projectID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, response{documents})
}

// GenerateCode handler to generate a new code for document
func (h HTTP) GenerateCode(c *gin.Context) {
	type uriVars struct {
		ID string `uri:"id" binding:"required"`
	}
	uriParams := uriVars{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		server.ErrorResp(c, err)
		return
	}
	documentID, err := strconv.Atoi(uriParams.ID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	document, err := h.service.GenerateCode(documentID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusOK, document)
}

// DeleteDocument handler to delete a document
func (h HTTP) DeleteDocument(c *gin.Context) {
	type uriVars struct {
		ID string `uri:"id" binding:"required"`
	}
	uriParams := uriVars{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		server.ErrorResp(c, err)
		return
	}
	documentID, err := strconv.Atoi(uriParams.ID)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	if err := h.service.Delete(documentID); err != nil {
		server.ErrorResp(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// ServeHTTP takes a gin instance and attach document handlers
func ServeHTTP(service Service, ginServer *gin.RouterGroup) {
	http := HTTP{service}
	ginServer.POST("/document", http.PostDocument)
	ginServer.GET("/project/:id/documents", http.ListDocuments)
	ginServer.PATCH("/document/:id/generate", http.GenerateCode)
	ginServer.DELETE("/document/:id", http.DeleteDocument)
}
