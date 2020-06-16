package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg"
	"github.com/vkevv/back-rectifier/pkg/config"
)

// New creates a new server
func New(conf config.Server) *gin.Engine {
	var server *gin.Engine
	if conf.Debug {
		server = gin.Default()
		gin.SetMode(gin.DebugMode)
		server.Use(gin.ErrorLogger())
	} else {
		server = gin.New()
		gin.SetMode(gin.ReleaseMode)
	}
	return server
}

// Start given a conf, starts the engine
func Start(conf config.Server, ge *gin.Engine) error {
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(conf.Port),
		Handler:      ge,
		ReadTimeout:  time.Duration(conf.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(conf.WriteTimeoutSeconds) * time.Second,
	}
	return server.ListenAndServe()
}

// ErrorResponse main struct to represent an error
type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	ErrStr  string `json:"error"`
}

// ErrorResp Generate error response
func ErrorResp(c *gin.Context, err error) {
	defer c.Abort()
	errorResp := ErrorResponse{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Unespected server error",
		ErrStr:  "",
	}
	if pkg.IsAPIError(err) {
		apiError := err.(pkg.APIError)
		errorResp.Code = apiError.Code
		errorResp.Message = apiError.Message
		if gin.Mode() == gin.DebugMode {
			errorResp.ErrStr = apiError.Err.Error()
		}
		c.JSON(http.StatusBadRequest, errorResp)
		return
	}
	if gin.Mode() == gin.DebugMode {
		errorResp.ErrStr = err.Error()
	}
	c.JSON(http.StatusInternalServerError, errorResp)
}
