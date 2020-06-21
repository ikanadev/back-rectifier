package server

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenParser represents JWT token parser
type TokenParser interface {
	ParseToken(string) (*jwt.Token, error)
}

// Auth handles token validation, it need the key to validate token
func Auth(key string, jwtToken TokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorResp := ErrorResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Unespected server error",
			ErrStr:  "",
		}
		// check if Authorization header exists, returns empty string if not
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" || tokenStr == "Bearer " || tokenStr == "Bearer" {
			c.Abort()
			errorResp.ErrStr = errors.New("Authorization header is missing").Error()
			c.JSON(http.StatusInternalServerError, errorResp)
			return
		}
		// TODO: verify expired token tokev.Valid
		token, err := jwtToken.ParseToken(tokenStr)
		if err != nil {
			errorResp.ErrStr = err.Error()
			c.Abort()
			c.JSON(http.StatusInternalServerError, errorResp)
			return
		}
		// cast token.Claims to jwt.MapClaims since token.Claims stores a map[string]interface{}
		data := token.Claims.(jwt.MapClaims)
		email, emailOk := data["email"]
		id, idOk := data["id"]
		// check if we can get values from map
		if !emailOk || !idOk {
			errorResp.ErrStr = errors.New("Can't get token claims").Error()
			c.Abort()
			c.JSON(http.StatusInternalServerError, errorResp)
			return
		}
		// Set context variables to use in handlers
		c.Set("email", email)
		c.Set("id", int(id.(float64)))
		c.Next()
	}
}

// CORSMiddleware to fix cors problem
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
