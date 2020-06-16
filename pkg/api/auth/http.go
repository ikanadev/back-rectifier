package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg/models"
	"github.com/vkevv/back-rectifier/pkg/server"
)

// HTTP is the main struct to handle all auth requests
type HTTP struct {
	service Service
}

// ServeHTTP takes a gin Engine and attach routes and handlers
func ServeHTTP(service Service, ginServer *gin.Engine, authMD gin.HandlerFunc) {
	http := HTTP{service: service}
	ginServer.POST("/login", http.Login)
	ginServer.POST("/register", http.Register)
	ginServer.GET("/token", authMD, http.ValidateToken)
}

// Login handler
func (h HTTP) Login(c *gin.Context) {
	type reqVars struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Token string      `json:"token"`
		User  models.User `json:"user"`
	}
	creds := reqVars{}
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	user, err := h.service.Login(creds.Email, creds.Password)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	tokenStr, err := h.service.GenerateToken(user)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	resp := response{
		Token: tokenStr,
		User:  user,
	}
	c.JSON(http.StatusOK, resp)
}

// Register handler
func (h HTTP) Register(c *gin.Context) {
	type reqVars struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
	}
	type response struct {
		Token string      `json:"token"`
		User  models.User `json:"user"`
	}
	creds := reqVars{}
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	user, err := h.service.Register(creds.Email, creds.Password, creds.Name, creds.LastName)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	tokenStr, err := h.service.GenerateToken(user)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	resp := response{
		Token: tokenStr,
		User:  user,
	}
	c.JSON(http.StatusOK, resp)
}

// ValidateToken validates a token
func (h HTTP) ValidateToken(c *gin.Context) {
	type response struct {
		Token string      `json:"token"`
		User  models.User `json:"user"`
	}
	user, err := h.service.Me(c)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	tokenStr, err := h.service.GenerateToken(user)
	if err != nil {
		server.ErrorResp(c, err)
		return
	}
	resp := response{
		Token: tokenStr,
		User:  user,
	}
	c.JSON(http.StatusOK, resp)
}
